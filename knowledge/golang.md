## 1、golang的gc算法，gc三色标记法，混合写屏障
**GC**：Garbage Collection（垃圾回收）





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

- work stealing机制
当本线程无可运行的G时，尝试从其他线程绑定的P中偷走G，而不是销毁线程
- hand off机制
当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行。
 抢占：在goroutine中要等待一个协程主动让出cpu才执行下一个协程，一个goroutine最多占用cpu 10ms，防止其他goroutine被饿死

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

## 4、go defer (for defer)

## 5、select可以用于什么？

常用于监听IO操作，当IO操作发生时，触发响应的动作，每个case语句必须是一个IO操作

1、所有channel表达式都会被求值，所有被发送的表达式都会被求值。求值顺序：自上而下，从左到右。。
结果是选择一个发送或接受的channel，无论选择哪一个case进行操作，表达式都会被执行。

2、如果有一个或多个IO操作可以完成，go运行时系统会随机选择一个执行，
否则，如果有default分支，执行default分支语句，如果没有default分支，select会一直阻塞，直到至少有一个IO操作可以执行。

## 6、context包的作用

## 7、

## 8、主协程如何等其余协程完再操作

1、使用channel通信

2、使用waitgroup

3、使用context

## 9、slice和array对比，slice扩容原理

？？ 同时通过append函数，slice底层数据结构是数组、len、cap组成。

## 10、map如何顺序读取

map不能顺序读取，是因为他是无序的，想要有序读取，就要将key变成有序的。
那么就需要将所有的key拿出来放到切片中排序，然后再取值

## 11、nil切片和空切片的区别

## 12、实现消息队列（多生产者，多消费者）

## 13、大文件排序

多路快排，然后再归并。

## 15、怎么做服务发现的

## 16、raft算法是那种一致性算法

## 17、raft有什么特点

## 18、

## 19、context的使用

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

## 24、内存逃逸

[内存逃逸](https://github.com/lifei6671/interview-go/blob/master/question/q019.md)

## 25、


## 26、源代码编译过程概述
[](https://draveness.me/golang/docs/part1-prerequisite/ch02-compile/golang-compile-intro/)


[可读](https://draveness.me/)