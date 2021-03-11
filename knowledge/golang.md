## 1、golang的内存分配

**内存分配**

![](https://gitee.com/zongl/cloudImage/raw/master/images/2021/03/02/go启动申请的虚拟内存.jpg)

go程序启动时向操作系统申请一块内存（虚拟的地址空间，并不会真正分配内存）

`arena`是堆，go动态分配的内存都在这个区域，它把内存分割成`8KB`大小的页，一些页组合起来成为`mspan`。

`bitmap`标识`arena`哪些地址保存了对象，并且用`4bit`标示位表示对象是否包含指针、GC信息。
`bitmap`中一个`byte`大小的内存对应`arena`区域中4个指针大小(指针大小为8B)的内存，所以`bitmap`大小为`512G/(4*8B)=16GB`。

`spans`存放`mspan`的指针，每个指针对应一页，所以`spans`的大小为`512GB/8KB*8B=512MB`

**内存管理单元**

`mspan`是golang中内存管理的基本单元，是一片连续的8KB的页组成的大块内存。
`mspan`是一个包含起始地址，mspan规格，页的数量等内容的双端链表。

**分配流程**

在栈上内存管理简单，分配比堆快。栈上的内存回收不需要程序关心，而堆需要GC。

变量是在栈上分配的还是在堆上分配的，是由逃逸分析的结果决定的。`go build -gcflags '-m -l' main.go`
内存分配器在分配对象时，会根据对象的大小进行分配：tiny对象（小于16B），小对象（大于16B，小于32KB），大对象（大于32KB）

mcache, mcentral, mheap是Go内存管理的三大组件，层层递进。
mcache管理M在本地缓存的mspan；mcentral管理全局的mspan供所有线程使用；mheap管理Go的所有动态分配内存。

- 32KB的对象在mheap分配
- `<=16B`的对象使用mcache的tiny分配器分配
- (16KB,32KB)的对象，根据对象的大小，使用相应规格的mspan分配
 * 如果mcache没有相应规格大小的mspan，向mcentral申请
 * 如果mcentral没有相应规格大小的mspan，向mheap申请
 * 如果mheap没有相应规格大小的mspan，向操作系统申请

[内存分配](https://www.cnblogs.com/shijingxiang/articles/11466957.html)
[内存分配](https://juejin.cn/post/6844903795739082760)

**v1.3之前的标记-清除（mark and sweep）算法

第一步，暂停程序业务逻辑，找出不可达的对象，然后做上标记；

第二步，回收标记好的对象。
但是，算法在执行过程中，会STW（Stop the world）

- 缺点：STW，程序出现卡顿；需要扫描整个heap；清理数据会产生heap碎片

**v1.5三色标记法**

![三色标记](https://gitee.com/zongl/cloudImage/raw/master/images/2021/03/02/三色标记法.gif)

1、只要是新创建的对象，默认都标记为“白色”

2、每次GC回收开始，从根节点（全局变量和 goroutine 栈对象）开始遍历所有对象，把遍历到的对象从白色集合放入“灰色”集合

3、遍历“灰色”集合，将灰色对象引用的对象从白色集合放入灰色集合，之后此灰色对象放入黑色对象

4、重复第三步，直至灰色中无任何对象

5、回收所有的白色标记对象，也就是回收垃圾

但是仍需要STW，因为在标记过程中，白色对象可以被黑色对象引用，他们之间的可达关系可能会遭到破坏

* 弱三色不等式：所有被黑色对象引用的白色对象都处于灰色保护状态
* 强三色不等式：不存在黑色到白色指针的对象

- 插入屏障：被黑色对象引用的对象强制变为灰色对象
- 删除屏障：被删除的对象，如果本身是灰色或白色，则被标记为灰色

**v1.8的混合写屏障（hybird write barrier）**

插入写屏障和删除写屏障的短板：
- 插入写屏障：结束时仍需要STW来重新扫描栈，标记栈上引用的白色对象存活
- 删除写屏障：回收精度低，GC开始时STW扫描堆栈来记录初始快照，这个过程会保护开始时刻的所有存活对象。

混合写屏障的规则

1、GC开始将栈上所有的对象全部扫描标记为黑色（之后不再进行第二次重复扫描，无需STW）

2、GC期间，任何在栈上创建的新对象，均为黑色

3、被删除的对象标记为灰色

4、被添加的对象标记为灰色

[GC垃圾](https://segmentfault.com/a/1190000022030353)
[](https://www.cnblogs.com/shijingxiang/articles/11466957.html)
[gc简明教程](https://zhuanlan.zhihu.com/p/92210761)

**何时触发GC**：会通过gcTrigger的test函数来决定是否需要触发GC

`/runtime/mgc.go gcTrigger.test()`
- `runtime.GC()`强制启动GC
- 再分配内存时，判断当前内存是否达到阈值会触发新一轮GC（比如当前为4MB，GOGC=100，4MB+4MB*GOGC/100）
- 上次GC间隔达到了runtime.forcegcperiod(默认2分钟)，会启动GC

**GC调优**

1、合理化内存分配速度

2、降低并复用已经申请的内存

3、调整GOGC

**内存逃逸**：变量通过了校验，可以在栈上分配；逃逸了，必须在堆上分配（编译期间逃逸分析）

- 在方法内把局部变量指针返回：局部变量原本应该在栈中分配，在栈中回收。但由于返回时被外部引用，因此生命周期大于栈，则溢出
- 发送指针或带有指针的值到channel
- 在一个切片存储指针或带指针的值
- 闭包引用逃逸
- slice的背后数组被重新分配了，因为append时超过了其容量（cap）。slice初始化是在栈上分配，但扩容时大于阈值会在堆上分配。
- 在interface类型上调用方法。在interface类型上调用方法都是动态调度的，方法的真正实现只能在运行时知道。


[](https://cjq99419.github.io/%E5%85%B3%E4%BA%8Ejava%E5%92%8Cgolang%E7%9A%84gc/)

## 2、go的调度
**GPM模型**

![](https://gitee.com/zongl/cloudImage/raw/master/images/2021/02/26/gpm1.jpeg)

- G：goroutine协程，每个goroutine都有自己独立的栈存放当前的运行内存及状态。可以把G看做是一个任务。
- P：processor处理器，可以认为一个“有运行任务”的P占用了一个cpu线程资源，且只要处于调度的时候就有P。
- M：线程，它本身和一个内核线程绑定

线程是运行goroutine的实体，调度器的功能是把可运行的goroutine分配到工作线程上。

golang在运行时存在两种队列，一种是全局队列，存放等待运行的goroutine；一种是P的本地队列，存量有限(不超过256个)。
新建G时，G优先加入到P的本地队列，如果队列满了，则会把本地队列中的一半G移动到全局队列中。
为了运行goroutine，M需要获取P，M会从P的本地队列中获取G，如果G队列为空时，M会尝试从全局队列中拿一批G放到P的本地队列或
从其他P的本地队列中**偷**一半放到自己P的本地队列中。

**P和M的数量问题**

- P的数量：有环境变量`GOMAXPROCS`或`runtime.GOMAXPROCS(int)`决定。这意味着程序在任意时刻都只有`GOMAXPROCS`个goroutine在同时运行。
- M的数量：go语言本身限制最大10,000个；一个M阻塞了，会创建新的M；`debug.SetMaxThreads()`设置M的最大数量

**P和M何时会被创建**

- P：在确定了P的最大数量后，运行时系统会根据这个数量创建P
- M：没有足够的M关联P并运行其中可运行的G。所有的M都被阻塞，而P有很多任务，会去找空闲的M，如果没有空闲的，就去创建新的M。

**调度器的设计策略**

复用线程：避免频繁的创建、销毁线程，而是对线程的复用。协程在用户态线程即完成切换，不会陷入到内核态，这种切换非常的轻量快速。

- work stealing机制
当本线程无可运行的G时，尝试从其他线程绑定的P中偷走G，而不是销毁线程
- hand off机制:
当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。
 + 利用并行：`GOMAXPROCS`设置p的数量，最多有`GOMAXPROCS`个线程分布在多个cpu同时运行。
 + 抢占：在goroutine中要等待一个协程主动让出cpu才执行下一个协程，一个goroutine最多占用cpu 10ms，防止其他goroutine被饿死

**go func()流程**
![](https://gitee.com/zongl/cloudImage/raw/master/images/2021/02/26/gpm2.jpeg)

1、通过go func() 来创建一个goroutine

2、有两个存储G的队列，一个是局部调度器P的本地队列，一个是全局G队列。新创建的G会先保存在P的本地队列中，如果本地队列满了会保存在全局队列中。

3、G只能运行在M中，一个M必须有一个P，M和P是1:1的关系。M会从P的本地队列弹出一个可执行的G来执行，如果P的本地队列为空，则会从其他P中偷取G来执行。

4、一个M调度G执行的过程是一个循环机制

5、当M执行一个G时发生系统调用、cgo调用或其他阻塞操作时，M会阻塞，如果当前有G在运行，系统会把线程M从P中摘除，然后创建一个新的M(有空闲的M就用空闲的M)来服务这个P。

6、当M系统调用结束时，这个G会尝试获取一个空闲的P执行，并放入P的本地队列。如果获取不到P，那么M变成休眠状态，加入到空闲线程中，G会放入全局队列中。

[Golang调度器GMP原理与调度](https://learnku.com/articles/41728)

## 3、go struct能不能比较

数组是可以比较的，切片、map，func是不能被比较的，如果结构体中包含这些就不能进行比较。

## 4、编译成汇编

`禁用优化和内敛进行编译 go tool compile -N -l -S stack.go > stack.s`

## 5、select可以用于什么？

常用于监听IO操作，当IO操作发生时，触发响应的动作，每个case语句必须是一个IO操作

1、所有channel表达式都会被求值，所有被发送的表达式都会被求值。求值顺序：自上而下，从左到右。。
结果是选择一个发送或接受的channel，无论选择哪一个case进行操作，表达式都会被执行。

2、如果有一个或多个IO操作可以完成，go运行时系统会随机选择一个执行，
否则，如果有default分支，执行default分支语句，如果没有default分支，select会一直阻塞，直到至少有一个IO操作可以执行。

## 6、context包的作用

```go
type Context interface {
    // 返回context被取消的时间，也是完成工作的截止时间
	Deadline() (deadline time.Time, ok bool)
    // 返回一个channel，这个channel会在当前工作完成或上下文被取消后关闭
	Done() <-chan struct{}
    // 返回context结束的原因，如果context被取消，会返回canceled；如果超时，会返回deadlineExceeded错误
	Err() error
    // 获取键对应的值
	Value(key interface{}) interface{}
}
```
context包通过构建树型关系的Context，来达到上一层Goroutine能对传递给下一层Goroutine的控制，在不同的goroutine之间对信号进行同步避免对计算资源的浪费
对于处理一个Request请求操作，需要采用context来层层控制Goroutine，以及传递一些变量来共享。
主要作用还是在多个goroutine组成的树中同步取消信号，以减少对资源的消耗和占用。

- Backgroupd(): 返回一个非nil，空的context，它不会被取消，没有值，也不过超时。通常用在main函数、初始化、和测试用例中，是顶级context
- TODO()：当不清楚context什么时候用，便于后期重构，先占个位。

![context的使用](https://segmentfault.com/a/1190000024441501)

## 7、client长连接

1、golang强制短连接

可以在请求头中加上`connection:close`,也可以设置`request`结构体`Close`变量为`true`。
```go
req,_ := http.NewRequest("Get","http://example.com",nil)
req.Close = true
```

2、长连接

golang client 默认开启，不设置主动断连，要保持长连接
 - 需要主动关闭`resp.Body.Close()`
 - 读完`body`里的数据，`ioutil.ReadAll(resp.body)`；`io.Copy(ioutil.Discard,resp.Body)
 
 可以通过`Transport`设置`DisableKeepAlives`关闭长连接，`MaxIdleConns`设置连接数，`MaxIdleConnsPerHost`控制连接host数

## 8、主协程如何等其余协程完再操作

1、使用channel通信

2、使用waitgroup

3、使用context

## 9、slice和array对比，slice扩容原理

nil切片和空切片的区别(分配空间)

同时通过append函数，slice底层数据结构是数组、len、cap组成。

**减少GC压力**

1、减少对象的分配

2、string和[]byte之间的转化，会对底层进行复制

3、少使用+连接string，通过append进行，最好能提前知道数组长度，进行空间预分配

## 10、map如何顺序读取

map不能顺序读取，是因为他是无序的，想要有序读取，就要将key变成有序的。
那么就需要将所有的key拿出来放到切片中排序，然后再取值

## 11、golang GDB调式



## 12、实现消息队列（多生产者，多消费者）

## 13、大文件排序

多路快排，然后再归并。

## 15、进程、线程、协程的区别

- 进程：进程具有一定独立功能的程序,有独立的内存空间，不同进程通过进程间通信来通信
- 线程：从属于进程，是进程的一个实例，是cpu调度和分派的基本单位，不拥有系统资源
- 协程：是用户态的轻量级线程，协程的调度完全由用户控制，有自己的寄存器和上下文

## 16、有缓冲通道是异步的，无缓冲通道是同步的

## 17、golang基础数据类型
- 数字类型：整型，浮点型，复数，int是32位还是64位看cpu
- 布尔型
- 字符串

## 18、

## 19、

## 20、

## 21、并发读写会发生什么？如果避免

**处理方法**

- 加锁
 * RWMutex
 * channel
- sync.map

## 22、uintptr和unsafe.Pointer的区别？
- unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型的指针，不可以参与指针的运算
- uintptr用于指针运算，GC不把uintptr当指针，也就是说uintptr无法持有对象，uintptr类型的目标会被回收
- unsafe.Pointer可以和普通指针进行相互转换
- unsafe.Pointer可以和uintptr进行相互转换

## 23、mutex的两种工作模式

```go
type Mutex struct {
	state int32 // 状态 , 第一位表示是否被锁，第二位表示是否唤醒，第3-32位表示等待的协程数量。1是锁，10是唤醒，11是饥饿
	sema  uint32
}
```
- 正常模式

所有等待锁的goroutine按照FIFO顺序等待。唤醒的goroutine不会直接拥有锁，而是会和新请求锁的goroutine竞争锁的拥有。
新请求锁的goroutine具有优势：它正在cpu上执行，可能有好几个，所以刚刚唤醒的goroutine有很大可能在锁竞争中失败。
在这种情况下，这个被唤醒的goroutine会加入到等待队列的前面。如果一个等待的goroutine超过1ms没有获取锁，那么它将会锁转变为饥饿状态。

- 饥饿模式

在饥饿模式下，锁的所有权将从unlock的goroutine直接交给等待队列的第一个。新来的goroutine将不会去尝试去获得锁，即时锁看起来是unlock状态，
也不会去尝试自旋操作(不挂起，继续占有cpu)，而是放在等待队列的尾部。

如果一个等待的goroutine获得了锁，并且满足其中任一条件
* 它是队列中最后一个
* 它等待的时间小于1ms

它会将锁的状态转换成正常状态。

正常状态有很好的性能表现，饥饿模式能够阻止尾部延迟的现象。

## 24、哈希表如何处理哈希冲突

[](https://www.cnblogs.com/zzdbullet/p/10512670.html)

## 25、


## 26、

[](https://www.cnblogs.com/wpgraceii/p/10528183.html)
[可读](https://draveness.me/)