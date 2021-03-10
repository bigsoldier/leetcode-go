## 1、gin的路由算法

## 2、beego的路由算法

## 3、如何调试golang的bug和性能问题？
- panic调用栈
- pprof
- 火焰图
- `go run -race`或`go build -race`来进行竞争检测
- 配合压测，查看系统磁盘IO/网络IO/内存占用/cpu占用

## 4、http 状态码
- 1**：信息，服务器收到请求，需要请求者继续执行操作
- 2**：成功，操作接受并处理
- 3**：重定向，需要进一步操作
- 4**：客户端错误
- 5**：服务器错误

## 5、遇到项目问题怎么处理

现象：我们的服务内存占用高，触发预警并自动重启，查看最近的提交，只有go版本升级和增加podman两个比较大的更新

猜想：
- 频繁重复申请对象资源，查看pprof，查看内存，cpu，goroutine的使用
- 不知名的内存泄露，top命令查看，VSZ(Virtual Memory Size)指标高
- prometheus的监控判别条件有问题，container_memory_usage_bytes = /container_memory_working_set_bytes = RSS+Cache
- go语言机制：
  * MADV_DONTNEED:内核会在进程的页表将这些页标记为"未分配"，从而进程的RSS就会尽快释放和变小。OS后续可以将对应的物理页分配给其他进程
  * MADV_FREE：内核只会在页表中将这些进程页面标记为可回收，在需要的时候才回收这些页面。(GO1.12-1.15)

方法：podman删除删除镜像

## 6、一些问题点

- sync.WaitGroup

```go
func main() {
	var wg sync.WaitGroup
	for i:=0;i < 5;i ++ {
		wg.Add(1)
		go func(wg sync.WaitGroup,i int) {
			defer wg.Done()
			fmt.Println(i)
		}(wg,i)
	}
	wg.Wait()
}
```
sync是不能被拷贝的，需要使用指针

## 7、项目实战

1、固有的瓶颈在jenkins这一块儿，通过测试jenkins的并发构建，发现前后端服务在并发量20左右，可能构建，在并发量30的时候会经常性的构建失败，
而且npm构建的速度比maven，go的构建慢很多，所以，第一步需要限制能够并发的数量，其次，根据需求有master，release，develop的优先级顺序，
所以我们对这些服务做了加权处理，用了小根堆来实现。

gitlab需要我们在构建服务之前打tag，由于nginx的work_processes有限制，所以同一时间不能并发构建，使用了线程池，

2、云平台，http和https的区别

3、运维平台，harbor接口升级兼容，我们做了一个通用的接口来兼容不同版本的harbor，虽然会增加一定的冗余，但能使接口更健壮

使用了helm的功能，对比优化了helm repo，不在容器内部做仓库管理，不做缓存，加快了安装的整体速度。

下载和安装功能使用条件变量来阻塞并发，协调响应访问共享资源的线程，直到有下载安装的请求，调用signal来唤醒阻塞线程，