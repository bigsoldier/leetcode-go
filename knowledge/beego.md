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