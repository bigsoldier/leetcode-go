## 1、CSP并发模型
[go语言并发模型](https://www.cnblogs.com/sunsky303/p/9115530.html)

## 2、go的调度

## 3、go struct能不能比较
数组是可以比较的，切片、map，func是不能被比较的，如果结构体中包含这些就不能进行比较。

## 4、go defer (for defer)

## 5、select可以用于什么？

## 6、context包的作用

## 7、client如何实现长连接

## 8、主协程如何等其余协程完再操作

## 9、slice、len，cap，共享，扩容

## 10、map如何顺序读取

## 11、实现set

## 12、实现消息队列（多生产者，多消费者）

## 13、大文件排序
多路快排，然后再归并。

## 15、怎么做服务发现的

## 16、raft算法是那种一致性算法

## 17、raft有什么特点

## 18、当go服务部署到线上了，发现有内存泄露，该怎么处理

## 19、context的使用

## 20、append的过程

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