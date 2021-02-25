## 1、为什么使用redis？
首先我们的项目在db上遇到了瓶颈，特别是批量查询和热点数据方面抗不住，需要缓存中间件的加入，目前的缓存中间件有redis和memcached。

[redis,mongodb,memcached的区别](https://www.cnblogs.com/tuyile006/p/6382062.html)

然后结合项目的特点，最后技术选型选择了redis。

## 2、redis跳表(skiplist=>zset)的原理？
给链表添加索引，可以添加多级索引，通过索引查找数据会简化查询。这样近似等于二分查找，使得查找的时间复杂度降到O(logn)。
为了避免增加和删除带来的索引调整问题，redis不要求上下相邻两层链表之间的个数有严格的对应关系，而是为每一层随机出一个层数(不超过当前层数+1)，最终层数比例大致为2:4:8:16...

而跳表的空间复杂度为O(n/2+n/4+n/8+...+4+2)=O(n-2)=O(n)

为什么不用B+树？因为跳表的时间复杂度和B+树一样，实现起来简单。

## 3、加入redis里有1亿个key，其中10w个key是以某个固定的前缀开头，如何把他们全部找出来？
使用keys命令可以扫描出指定模式的key列表，但由于redis是单线程的，keys指令会导致线程阻塞一段时间，线上服务会停顿，直到指令执行完毕，服务才能恢复。
可以使用scan指令可以无阻塞去除指定模式的key列表，但可能会有重复，需要在客户端去重，整体花费时间大于keys。

## 4、redis缓存穿透，击穿，雪崩
**雪崩**
1、原因
统一时间缓存大面积失效，所有的请求发送到数据库，db扛不住
2、处理方法
批量往redis存储数据时，把每个key的失效时间加上个随机数，可以保证数据不会再同一时间大面积失效

**穿透**
1、原因
用户不断发起请求，在缓存和db中都没有，造成db压力过大，每次查询都绕开缓存查db
2、处理方法
- 接口层添加校验，不合法的参数直接返回
- 在缓存查不到，db也没有的情况，可以将对应的key的value写为null，或其他特殊值，同时将过期时间设置短一点，以免影响正常情况，可以防止反复使用同一请求暴力攻击
- 正常用户不会暴力攻击，可以为每一个ip设置访问阈值
- redis的高级用法布隆过滤器(Boolm Filter),利用高效的数据结构和算法判断key是否在数据库存在，不存在直接return，存在去查询db再刷新kv再返回

*布隆过滤器*
布隆需要记录见过的数据，这里的记录需要通过hash函数对数据进行hash，得到数组下标并存储在bit数组里标记为1，这样记录一个数据只需要1比特。
给布隆过滤器一个数据，进行hash得到下标，如果从bit数组中取得数据是1则数据可能存在，如果是0，肯定不存在。

由于hash算法存在碰撞可能性，所以不同数据可能hash为同一个下标，故为提高精确，使用多个hash算法标记一个数据和增大bit数组的大小。

**击穿**
1、原因
一个key是热点数据，不停的抗大并发请求，全部集中访问，当key过期的瞬间，大并发击穿缓存，直接请求数据库。
2、处理方法
- 设置热点数据永不过期
- 加上互斥锁

## 5、redis的数据结构
1、字符串string
使用动态字符串(SDS)的抽象类型
```cgo
struct sdshdr {
    int len;    // buf已占用的空间长度
    int free;   // buf中剩余的空间长度
    char buf[]; // 数据空间
}
```
为什么不用C字符串呢？
- 计数方式不同：
C语言对于字符串长度的统计完全来自遍历，从头遍历到末尾，发现空字符停止，时间复杂度为O(n);
SDS可以直接获取字符串长度，时间复杂度为O(1).
- 杜绝缓冲区溢出
C是不记录字符串长度的，如果我们调用了拼接的函数，如果没有提前计算好内存，是会产生缓存区溢出的；
SDS增加字符串长度需要验证free的长度，如果free不够会扩容整个buf，防止溢出。
- 减少修改字符串时带来的内存重分配次数
redis作为高速缓存数据库，需要对字符串进行频繁的修改，采用两种方法性能最大化：
  * 空间预分配：对SDS进行扩展操作时，会为SDS分配好内存，如果长度小于1MB那len=2*len+1,如果长度大于1MB那len=1MB+len+1
  * 惰性空间释放：如果缩短SDS的字符串长度，redis并不是马上减少SDS所占用内存，只是增加free长度，同时向外提供api，真正需要释放的时候，才去重新缩小SDS所占内存
- 二进制安全
C语言字符串是以'\0'作为字符串的结束标记，而SDS是使用len的长度来标记字符串的结束，所以SDS可以存储字符串以外的任意二进制流，也就是说SDS不依赖'\0'为结束的依据。

[动态字符串](https://segmentfault.com/a/1190000023054174)

- 字典hash
- 列表list
- 集合set
- 有序集合sorted set

**场景**
- string：缓存功能；计数器；共享session
- hash：类似与map
- list：存储一些列表性数据，如消息队列；分页展示数据；
- set：用于某个系统部署在多台机器上，使用全局去重
- sorted set：排行榜；权重任务

**中级用法**
- hyperloglog：基数统计的算法，可以用来做大规模的数据去重计数功能
- geo：存储地理位置信息
- pub/sub：发布订阅功能，可用作简单的消息队列

**高级用法**
- boolmFilter：布隆过滤器，在4中有介绍
- pipeline：可以批量执行一组指令，一次性返回全部结果，可以减少频繁的请求应答

## 6、redis对比memcached有哪些优势？
- memcached所有值均是简单字符串，redis有更丰富的数据结构
- redis的速度比memcached快很多
- redis可以持久化数据

## 7、redis的持久化方案
**全量持久化RDB** (`dump.rdb`)

按照一定的时间周期将目前服务中所有的数据全部写入磁盘。
在操作过程中，主线程会fork一个线程专门用来拷贝，
而这段时间的数据变化会以副本的方式存放在一个内存区域，在快照操作完成后同步到原来的内存区域。

*命令*
手动执行 save/bgsave

或配置 save <seconds> <changes> (在多少秒内有多少key信息发生变化，则进行快照)

**优势**
- rdb是redis数据非常紧凑的单文件时间点，非常适合备份
- rdb对于灾难恢复非常有用
- rdb最大限度提高redis的性能，父进程不会执行磁盘I/O操作
- 与aof相比，rdb允许大型数据集更快重启

**缺点**
- 在redis处于任何原因没有在正确关闭下停止工作，会丢失数据
- rdb需要经常fork子进程，如果数据集很大，fork操作会很耗时，如果数据集大并且cpu性能不佳，会导致redis停止

**增量持久化AOF**

配置文件中 `appendonly yes`

每次redis收到更改数据集时都会添加到aof。每次重启redis时，都会重新播放aof以重建状态。

可以猜想到，随着写操作越来越多，aof越来越大，redis为aof提供重写aof功能，保证aof可以存储尽可能少的操作命令就能保证数据恢复到最新状态。

**优势**
- 更加持久：可以采用不同的fsync策略：always(每个写操作)、everysec(每秒，默认值)、no。
- aof是追加的数据，如果处于某种问题redis停止工作，不会出现数据丢失情况。
- 当aof文件太大时，redis能够在后台重写aof
- aof以易于理解和解析的格式包含所有操作，可以轻松导出aof文件

**缺点**
- 对于相同的数据集，aof文件通常大于rdb文件
- 根据fsync策略，aof可能比rdb慢

## 8、使用redis实现分布式锁
**加锁**
如果使用SETNX加锁会出现过期后key删除但没有删除锁。

`SET lock_key random_value NX PX 5000`

`random_key`是客户端生成的唯一字符串

`NX`表示只在键不存在时，对键进行设置操作

`PX 5000`设置键的过期时间为5000毫秒

**解锁**
删除key，通常需要先比较value，避免将其他客户端的锁给删除了

## 9、redis的异步队列
使用list结构作为队列，rpush生产消息，lpop消费消息。当lpop没有消息时，可以sleep后重试。

不sleep的方式：使用blpop，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。

能否生产一次消费多次？使用pub/sub主题订阅者模式，实现1：N的消息队列。

发布/订阅的缺陷？客户端在执行订阅操作的过程中断网，客户端会丢失断线期间的消息，可以使用专业的消息队列。

延迟队列？使用sortedset，拿时间戳做score，消息内容作为key调用zadd来生产消息，zrangebyscore指令获取N秒之前的数据轮询处理。

## 10、如果有大量的key需要设置同一时间过期，需要注意什么？
如果大量的key过期时间设置的过于集中，到过期的那个时间点，redis可能会出现短暂的卡顿现象，一般需要在时间上加一个随机值，使过期时间分散。

## 11、pipeline有什么好处，为什么要使用pipeline？
可以将多次IO往返的时间缩减为一次，前提是pipeline执行的指令之间没有因果关系。
使用redis-benchmark进行压测的时候可以发现影响redis的QPS峰值的一个重要因素是pipeline批次指令的数目。

## 12、redis同步机制
**全同步**：slave启动时进行的初始化同步

- 在slave启动时，会向master发送一条sync指令：psync ? -1,格式为psync {runId} {offset}
- master收到指令后，会启动一个备份进程将所有数据写到rdb文件中去
- 更新master状态(备份是否成功、备份时间等)，然后将rdb文件内容发送给等待中的slave

**部分同步**： redis运行过程中的修改同步

当redis的master/slave服务启动后，首先进行全同步。之后，所有的写操作都在master上，而所有的读操作都在slave上。因此写操作需要及时同步到所有的slave上。
- master收到一个操作，然后判断是否需要同步到slave, psync {runId} {offset}
- 如果需要同步，则将操作记录到aof中
- 遍历所有的slave，将操作的指令和参数写入到slave的回复缓存中
- 一旦slave对应的socket发送缓存中有空间写入数据，即可将数据通过socket发送出去

**注意事项**
- 复制超时：针对数据量大的节点，设置超时时间`repl-timeout`(默认60s)，防止出现全量同步数据超时
- 复制积压缓冲区溢出：slave在开始接受rdb到接受完毕期间，主节点仍然响应读写命令，因此主节点会把这期间的写入命令保存在赋值积压缓冲区，当从节点加载完rdb文件后，主节点在把缓冲区的数据发送给从节点，保证主从数据一致。
如果主节点创建和传输rdb时间过长，会导致主节点复制客户端缓冲区溢出，默认为`client-output-buffer-limit slave 256MB 64MB 60`,如果60秒内缓冲区消耗持续大于64MB或直接超过256MB，主节点将直接关闭复制客户端连接，造成全量同步失败。
- slave全量同步时的响应问题：slave节点在接受完主节点传送的全部数据后会清空自身旧数据，然后加载rdb，对于较大的rdb，这一步依然比较耗时。
对于线上做读写分离的场景，从节点负责响应读命令，如果slave节点处于全量复制阶段，那么slave节点在响应读命令时可能拿到过期或错误的数据。对于这一问题，redis提供参数控制是否关闭保证一致性。

[redis数据同步原理](https://my.oschina.net/u/585635/blog/3220828)

## 13、redis集群及高可用？
1、redis单副本
采用单个redis节点部署，没有备用节点实时同步数据，不提供数据持久化和备份策略，适用于数据可靠性要求不高的纯缓存业务场景。

**优点**
- 架构简单，部署方便，高性价比，高性能

**缺点**
- 不保证数据的可靠性
- 在缓存中使用，进程重启后数据丢失，即时有备用节点解决高可用性，但仍不能解决缓存预热问题
- 高性能受限于单核cpu的处理能力，cpu成为主要瓶颈

2、redis多副本(主从)
相较于单副本最大的特点就是主从实例间数据实时同步，并且提供数据持久化和备份策略。

**优点**
- 高可靠性：采用双机主备架构，能在主库出现故障时自动进行主被切换，从库提升为主库提供服务，保证服务平稳运行；
另一方面，开启数据持久化功能和配置合理的备份策略，能有效解决数据误操作和数据异常丢失问题。
- 读写分离策略：从节点可以扩展主库的写能力，有效应对大并发的读操作

**缺点**
- 故障恢复复发，当主库节点出现故障，需要手动将一个从节点晋升为主节点，同时通知业务方变更配置，并且需要让其他从库节点去赋值新主库节点，需要人为干预
- 主库写能力受到单机限制，可以考虑分片
- 主库存储受到单机限制，可以考虑pika

3、redis sentinel(哨兵)
可以实现故障发现，故障自动转移，配置中心和客户端通知，节点数需要满足奇数个

**优点**
- 部署简单
- 能够解决redis主从模式下的高可用切换问题
- 很方便实现redis数据节点的线性扩展，轻松突破redis自身单线程瓶颈，可极大满足redis大容量或高性能的业务需求
- 可以实现一套sentinel监控一组redis数据节点或多组数据节点

**缺点**
- 部署相对于redis主从复杂，原理理解繁琐
- 资源浪费，redis数据节点中slave节点作为备份节点不提供服务
- 不能解决读写分离问题

4、redis cluster
分布式集群方案，所欲节点通过服务通道直接相连，各个节点之间通过二进制协议优化传输速度和带宽。客户端直连其中一个节点，进行ascii协议通信

**优点**
- 无中心架构
- 数据按slot存储在多个节点，数据间数据共享，可动态调整数据分布
- 可线性扩展到1000多个节点，节点可动态添加和删除
- 高可用，通过增加slave做standby数据副本，能够实现故障自动failover
- 降低运维成本，提高系统的扩展性和可用性

**缺点**
- client实现复杂
- 节点会因为某些原因发生阻塞(阻塞时间大于`cluster-node-timeout`)，被判断下线
- 数据通过异步复制，不保证数据的强一致性
- 多个业务使用同一套集群时，无法根据统计区分冷热数据，资源隔离性差，容易相互影响
- slave在集群充当冷备，不能缓解读压力

5、redis自研高可用框架
主要体现在配置中心、故障探测和 failover 的处理机制上，通常需要根据企业业务的实际线上环境来定制化。

**优点**
- 高可靠性，高可用性
- 自主可控
- 贴切业务实际需求，兼容性好

**缺点**
- 实现复杂，开发成本高
- 需要建立配套设置，如监控，域名，存储等
- 维护成本高

6、redis代理中间件
**codis**

能够把所有的redis实例当成一个来使用，因为codis是无状态的，可以增加多个codis来提升qps，同时能起到容灾的作用。

[redis高可用解决方案总结](https://www.jianshu.com/p/5de2ab291696)

哨兵着重于高可用，在master宕机后自动将slave提升为master；cluster着重于可扩展性，在单个redis内存不足时，使用cluster分片存储。

## 14、redis为什么这么快？
- 完全基于内存，数据存在内存中，类似于hashmap，查找和操作的时间复杂度都是O(1)
- 高效的数据结构
- 采用单线程，避免了不必要的上下文切换和竞争条件，也不存在多线程或多线程导致的切换而消耗cpu，不用考虑各种锁的问题
- 采用多路I/O复用模型，非阻塞IO
- 删除过期数据的不同策略
  * 定时删除：在键过期时间的同时，创建一个定时器，让定时器在键过期时来删除键。对cpu不友好，会占用cpu性能
  * 惰性删除：键过期后不管，每次读取键时判断键是否过期，如果过期删除键返回空。对内存不友好，有的键过期但没有访问不会被删除
  * 定期删除：每隔一段时间对数据库中过期的键进行一次检查。两种方式的折中

## 15、redis几种数据淘汰策略？
- noeviction: 返回错误当内存限制达到并且客户端尝试执行会让更多内存被使用的命令
- allkeys-lru: 尝试回收最少使用的键，使得新添加的数据有空间存放
- volatile-lru: 尝试回收最少使用的键，但仅限于在过期集合的键，使得新添加的数据有空间存放
- allkeys-random：回收随机的键使得新添加的数据有空间存放
- volatile-random：回收随机的键使得新添加的数据有空间存放，但仅限于在过期集合的键
- volatile-ttl：回收在过期集合的键，优先回收存活时间ttl较短的键，使得新添加的数据有空间存放

**LRU**(最近最少使用):算法实现(哈希表+双向链表)
```go
package main

type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head,tail *DLinkedNode
}

// 双向链表
type DLinkedNode struct {
	key,value int
	prev,next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key: key,value: value}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache: map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key,value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache,removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
```

## 16、kv，db读写模式？
- 读的时候，先读缓存，缓存没有的话读数据库，然后取出数据放入缓存，同时返回响应
- 写的时候，先更新数据库，然后再删除缓存

## 17、redis事务
事务本质是一组命令的集合，事务一次性执行多个命令，一个事务所有命令都会被序列化，
在事务执行过程中，会按照顺序串行执行队列中的命令，其他客户端的请求不会插入到事务执行中。

事务不保证原子性，且没有回滚。任意命令执行失败，其余命令仍会执行。

```text
multi 标记一个事务开启
exec 执行事务块内的命令
discard 取消事务，放弃事务块中的所有命令
watch 监视一个(或多个)key，如果在事务执行之前key被其他命令改动，那么事务被打断
unwatch 取消对所有key的监视
```

## 18、redis设置密码及验证密码
设置密码： config set requirepass "password"

验证密码： auth "password"

## 19、redis一致性hash
**数据分片**

分布式存储数据，经常考虑数据分片，避免将大量的数据存储在单表或单库中，造成查询时间过长。
比如有3个库，存储的数据 hash(id)%3 来找到对应的库。

但是会存在一个问题，当机器不够需要扩容或机器宕机时，机器数量发生变化，造成数据命中率下降。

**一致性hash**

是对2^32取模，简单来说，一致性hash算法将整个哈希值空间组成一个虚拟的圆环，如假设某哈希函数的值空间是 [0,2^32-1]。

下一步将各个服务器使用服务器的ip或主机名作为关键字进行哈希，这样每台机器都能确定在哈希环的位置。

将数据key使用相同的哈希函数计算出哈希值，并确定数据在环上的位置，将数据从位置顺指针找到第一台的服务器节点，这个节点就是key存储的服务器。

![一致性哈希](https://gitee.com/zongl/cloudImage/raw/master/images/2021/02/25/一致性哈希.webp)

**容错性和可扩展性**
- 单个系统宕机，仅会影响节点前的key对象，不会影响其他系统
- 新增节点，仅会影响新节点到上个节点之间的key对象，不会影响其他数据

**数据倾斜**

如果一致性哈希算法在服务器节点太少时，容器因为节点分布不均匀而造成数据倾斜(被缓存的对象大部分集中缓存在某一台服务器上)。

为解决这个问题，引入虚拟节点机制，即对每一个服务节点计算多个哈希值，每个计算结果都放置一个服务节点，称为虚拟节点。
具体做法可以在主机名后面添加编号来实现，如"node01#1","node01#2"等，这样很少的服务节点也能做到相对均匀的数据分布。

[redis一致性哈希](https://juejin.cn/post/6850418113830846471)