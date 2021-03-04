## docker 核心原理

**限制的视图namespace**

| namespace | 隔离内容                   |
| --------- | -------------------------- |
| IPC       | 信号量、消息队列和共享内存 |
| Network   | 网络设备、网络栈、端口等   |
| Mount     | 挂载点(文件系统)           |
| PID       | 进程编号                   |
| User      | 用户和用户组               |
| UTC       | 主机名和域名               |

**限制资源cgroups：包括cpu、内存、磁盘、网络带宽等**

- cpu
- blkio，为块设备设定I/O限制，一般用于磁盘
- cpuset，为进程分配独立的cpu核和对应的内存节点
- memory，内存

**rootfs**

`chroot`:change root file system。改变进程的根目录到你指定的目录

1、启动linux namespace配置
2、设置指定的cgroups参数
3、切换进程的根目录

## 2、pivot_root和chroot的区别
pivot_root主要是把整个系统切换到一个新的root目录，而移除对之前root文件系统的依赖，这样你能够umount原先的root文件系统。
chroot是针对进程，而系统的其他部分仍运行在老的root目录

## 3、Dockerfile内容
- FROM：指定基础镜像
- MAINTAINER：维护者信息
- RUN：构建镜像时执行的命令
- ADD：将本地文件添加到容器中，tar类型文件会自动解压(网络压缩资源不能解压)
- COPY：类似于ADD，但不能自动解压文件，也不能访问网络资源
- CMD：容器启动后调用，在`docker run`时运行
- ENTRYPOINT：类似于CMD，但不会被`docker run`的命令行参数的指令覆盖
- LABLE：为镜像添加元数据
- ENV：环境变量
- ARG：构建参数，仅在`docker build`有效，构建好的镜像不存在此环境变量
- VOLUME：持久化目录
= EXPOSE：外界交互端口
- WORKDIR：工作目录
- USER：执行后续命令的用户和用户组
- HEALTHCHECK：健康检查
- ONBUILD：延迟构建命令，有新的Dockerfile使用之前的镜像会执行ONBUILD命令

## 4、优化Dockerfile
- 构建顺序影响缓存利用率。
把不需要经常更改的行放在最前面，更改频繁的行放在最后面
- 只拷贝文件，防止溢出
- 最小化可缓存的执行层，每一个`RUN`指令都被看作可缓存的执行单元，太多的`RUN`指令会增加镜像层数，增大体积。
将更新缓存和安装文件放在同一个`RUN`指令中。
- 减小镜像体积：删除不必要的依赖 `apt --no-install-recommends`；删除包管理工具缓存
- 使用多阶段构建

## 5、docker load 加载一个镜像，docker images查看不到的原因


## 6、docker后端存储


[docker 面试](https://www.jianshu.com/p/2de643caefc1)

## 7、docker预热

没有平台之前，运维人员安装微服务的时候，需要先拉取镜像传到内网，然后再执行安装