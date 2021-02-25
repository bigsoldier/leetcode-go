## 1、k8s pause容器的作用？
pause全称是infra容器，在pod启动时，pause容器是第一个被创建，而其他用户定义的容器，则通过join network namespace的方式，与infra容器关联在一起。
这个镜像是用汇编语言编写的，永远处于"暂停"状态的容器。

这样对于pod里的容器来说：
- 它们可以直接通过localhost相互通信
- 它们看到的网络设备和infra容器看到的网络设备完全一样
- 一个pod只有一个ip，也就是这个pod的network namespace对于的ip地址
- infra的生命周期只与pod的生命生命周期一致