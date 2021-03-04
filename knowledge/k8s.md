## 1、k8s pause容器的作用？
pause全称是infra容器，在pod启动时，pause容器是第一个被创建，而其他用户定义的容器，则通过join network namespace的方式，与infra容器关联在一起。
这个镜像是用汇编语言编写的，永远处于"暂停"状态的容器。

这样对于pod里的容器来说：
- 它们可以直接通过localhost相互通信
- 它们看到的网络设备和infra容器看到的网络设备完全一样
- 一个pod只有一个ip，也就是这个pod的network namespace对于的ip地址
- infra的生命周期只与pod的生命生命周期一致

## 2、k8s如何调度一个pod

**创建一个pod**

1、向kube-apiserver发出指令

2、api响应命令，通过一系列认证授权，将pod数据存储到etcd，创建deployment资源并初始化

3、controller通过list-watch机制，监测发现新的deployment，将该资源加入到内部工作队列中，发现该资源没有关联的pod和replicaset，
启用deployment controller创建replicaset，再启用replicaset controller创建pod

4、所有controller被创建完成后，将deployment，replicaset，pod资源更新到etcd

5、scheduler通过list-watch机制，监测发现新的pod，经过主机过滤、主机打分机制，将pod绑定到合适的主机

6、将绑定结果更新到etcd

7、kubelet每隔20s向apiserver通过nodename获取自身node上所要运行的pod清单，通过与自己内部缓存进行比较，新增pod

8、kubelet创建pod

9、kube-proxy为新创建的pod注册动态dns到coreOS，给pod的servi添加iptables/ipvs规则，用于服务发现和负载均衡

10、controller通过control loop将当前pod状态与用户所期望的状态做对比，如果当前状态与用户期望状态不同，则controller会将pod修改为用户期望状态，否则会删除pod，重新创建pod。

[k8s创建pod](https://developer.51cto.com/art/202010/630004.htm?pc)

**pod调度流程**

1、预选阶段：会将所有满足pod调度需求的node选出来
 - node上是否存在当前被调度pod的端口
 - 是否在pod定义了nodeSelector，绑定了nodeName
 - 检查node是否有空闲资源
 - 检查pod请求的卷在node上是否可用
 - 根据资源（内存，进程id，磁盘，cpu）进行调度
 - 检查pod的容忍度是否被打污点

2、打分阶段：在过滤阶段后调度器会为pod从所有可调度节点中选取一个最合适的node
 - 优先减少同一个service或rc的pod数量
 - 优先将pod调度到相同的拓扑上（如同一个节点等）
 - 优先调度到pod少及资源使用少的节点
 - 尽量调度到已经使用过的node上
 - 优先平衡各节点的资源使用
 - 优先调度到nodeAffinity（node亲和度）和TaintToleration（污点）节点上
 - 尽量使用已经拉过镜像的节点上

[](https://cloud.tencent.com/developer/article/1644857)

## 3、pod的生命周期

pod的生命周期的变化，主要体现在pod api对象的status部分。其中`pod.status.phase`就是pod的状态。

**pod状态**
- Pending：请求创建pod时，条件不满足，调度没有完成，没有节点适合满足调度条件
- Running：pod已经调度成功，跟具体的节点绑定。它包含的容器都已经创建成功，并且至少有一个正在运行
- Failed：pod里至少有一个容器以不正常状态（状态码不为0）退出。出现这种状态需要debug容器，比如查看event和日志
- Succeeded：pod里所有的容器都正常运行完毕，并且已经退出了
- Unknown：异常状态，意味着pod的状态不能持续地被kubelet汇报给kube-apiserver，这很可能是主从节点间的通信出现了问题。

Ready这个状态标识pod已经正常启动，并且可以对外提供服务。

## 4、健康检查存活和就绪探针有什么区别

- livenessProbe(存活指针)：如果成功，说明pod的状态为running
- readinessProbe（就绪指针）：

## 5、deployment、statefulset有什么区别

## 6、crd和operator

## 7、cni？k8s集群使用的网络插件
flannel

## 8、为什么pod资源有request和limit

## 9、一个请求到pod接受响应，经过哪些过程

## 10、prometheus监控架构？采集数据流向

查询节点的数据来源于cadvisor

![](https://prometheus.io/assets/architecture.png)

prometheus从exporter或pushgateway拉取数据（如果在k8s内部用服务发现机制），
默认本地存储抓取的数据，并通过一定的规则清理和整理数据，并把得到的数据存储在新的时间序列中，采集到的数据有两个方向，一个是报警Alertmanager，一个是可视化grafana。
promQL和其他可视化api可视化的展示收集的数据，并通过Alertmanager提供报警能力。

- Prometheus Server负责从Exporter拉取和存储监控数据，并提供查询语言PromQL
  * Retrival：采样模块
  * TSDB：存储模块默认本地存储为tsdb
  * http server：提供http接口查询和面板
- exporter/jobs：负责收集目标对象的性能数据，并通过http接口供prometheus server获取。
- pushgateway：主要用于短期jobs，由于这类jobs存在时间比较短，可能在prometheus来pull之前就消失了。
- Alertmanager：接受到alerts后，会进行去除重复，分组，发出告警。
- service discovery服务发现：支持多种服务发现机制：文件、dns、k8s

**存储**

TSDB时序数据库，将监控数据以时间段翟芬不同的block，并且会压缩合并历史数据；引入wal，避免宕机导致的数据丢失问题。

block包括chunk，index，meta.json，tombstones。

随着数据量的增长，tsdb会将小的block合并成大的block，这样不仅可以减少数据存储，还可以减少内存中block个数，便于对数据进行检索。

每个block都有一个全局的名称。block的命名规则：时间戳（6个字节）+ 随机数（10个字节）。prometheus将16字节的名称通过算法转化为26字节可排序的字符串。
可以通过block文件名确定block的创建时间，可以很方便按时间对block进行排序。

- chunk：用于保存压缩后的数据
- index：用于对监控数据进行快读的检索，记录chunk的偏移位置
- tombstone：tsdb在删除block数据块时会将整个目录删除，但如果只删除一部分数据块时，可通过tombstone进行软删除
- meta.json：记录block的元数据信息，包括一个数据块记录样本的起始时间，截止时间，样本数等信息。

- wal：预写日志，方便数据回滚、重试等操作保证数据的可靠性。为了防止暂存在内存中还未写入磁盘的监控数据

[prometheus存储层演进](https://xie.infoq.cn/article/9071f261190acbdf73dfcf4d7)

## 11、k8s整体架构

**功能**
- 自动化容器部署和复制
- 实时弹性伸缩容器规模
- 容器编排成组，并提供容器建的负载均衡

**核心组件**

- etcd：保存了集群的状态
- apiserver：提供资源操作的唯一入口，并提供认证、授权、访问控制、api注册和发现等机制
- controller manager负责维护集群的状态，比如故障检测、自动扩展、滚动更新等
- scheduler负责资源的调度，按照预定的调度策略将pod调度到相应的机器上
- kubelet：负责维护容器的生命周期，同时也负责volume和cni的管理
- container runtime：负责镜像管理以及pod和容器的真正运行（cni）
- kube-proxy：负责为service提供cluster内部的服务发现和负载均衡



## 12、service和ingress

## 13、etcd的raft

[](https://zhuanlan.zhihu.com/p/91288179)
