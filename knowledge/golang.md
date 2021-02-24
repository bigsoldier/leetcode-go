### 1、CSP并发模型
[go语言并发模型](https://www.cnblogs.com/sunsky303/p/9115530.html)

### 2、go的调度

### 3、go struct能不能比较
数组是可以比较的，切片、map，func是不能被比较的，如果结构体中包含这些就不能进行比较。

### 4、go defer (for defer)

### 5、select可以用于什么？

### 6、context包的作用

### 7、client如何实现长连接

### 8、主协程如何等其余协程完再操作

### 9、slice、len，cap，共享，扩容

### 10、map如何顺序读取

### 11、实现set

### 12、实现消息队列（多生产者，多消费者）

### 13、大文件排序
多路快排，然后再归并。

### 15、怎么做服务发现的

### 16、raft算法是那种一致性算法

### 17、raft有什么特点

### 18、当go服务部署到线上了，发现有内存泄露，该怎么处理

### 19、context的使用

### 20、append的过程

### 21、并发读写会发生什么？如果避免

### 22、uintptr和unsafe.Pointer的区别？
- unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型的指针，不可以参与指针的运算
- uintptr用于指针运算，GC不把uintptr当指针，也就是说uintptr无法持有对象，uintptr类型的目标会被回收
- unsafe.Pointer可以和普通指针进行相互转换
- unsafe.Pointer可以和uintptr进行相互转换



