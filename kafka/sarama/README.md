https://cloud.tencent.com/developer/article/1823421

http://www.catddm.com/archives/773

https://github.com/Shopify/sarama

https://blog.csdn.net/justlpf/article/details/107400534

https://www.cnblogs.com/zhaoxianxin/p/13432068.html
https://blog.csdn.net/luslin1711/article/details/105798571/


2021/09/11 22:21:54 kafka server: Replication-factor is invalid.
warn [producer clientid=console-producer] error while fetching metadata with correlation id 140 : {test=
warn [controller id=0, targetbrokerid=0] connection to node 0 (/202.112.238.170:9092) could not be established. broker may not be available.

 % vi config/server.properties
 listeners=PLAINTEXT://:9092
 advertised.listeners=PLAINTEXT://localhost:9092



https://help.aliyun.com/document_detail/266782.html

Sarama Go客户端存在以下已知问题：

当Topic新增分区时，Sarama Go客户端无法感知并消费新增分区，需要客户端重启后，才能消费到新增分区。
当Sarama Go客户端同时订阅两个以上的Topic时，有可能会导致部分分区无法正常消费消息。
当Sarama Go客户端的消费位点重置策略设置为Oldest(earliest)时，如果客户端宕机或服务端版本升级，由于Sarama Go客户端自行实现OutOfRange机制，有可能会导致客户端从最小位点开始重新消费所有消息。
解决方案
建议尽早将Sarama Go客户端替换为Confluent Go客户端。

https://zhuanlan.zhihu.com/p/110114004


Group Coordinator
在0.9版本之前，Kafka强依赖于ZooKeeper实现Consumer Group的管理：

Group内每个Consumer通过在ZK内抢注节点来决定消费哪些Partition，并注册对Group和Broker相关节点的监听，以获知消费环境的变化（其他Consumer掉线、Broker宕机等），进而触发Rebalance；
Offset值也维护在ZK中，老生常谈了。
这种方式除了过于依赖ZK，导致ZK压力偏大之外，还有两个分布式系统中常见且严重的问题：

羊群效应（herd effect）——一个被监听的ZK节点发生变化，导致大量的通知发送给所有监听者（即Consumer）；
脑裂（split brain）——ZK只保证最终一致性，不同的Consumer在同一时刻可能看到不同的Group和Broker状态，造成Rebalance混乱。
所以从0.9版本开始，Kafka重新设计了名为Group Coordinator的“协调者”服务负责实现上述职责，将这部分工作从ZK剥离开来。每个Broker在启动时，都会顺带启动一个Group Coordinator实例。每个Consumer Group在初始化时，都会分配给一个Group Coordinator实例来管理消费关系和Offset
Group Coordinator提交Offset时也不再是向ZK写，而是写入那个广为人知的特殊Topic——__consumer_offsets里。key是group-topic-partition格式的，value为Offset值。

那么该如何确定一个Consumer Group被分配给哪个Group Coordinator呢？Kafka根据groupId.hashCode() % offsets.topic.num.partitions取绝对值来得出该Group的Offset信息写入__consumer_offsets的分区号，并将Group分配给该分区Leader所在的Broker上的那个Group Coordinator。

整个Rebalance分为两个大步骤：JOIN和SYNC。
所有Consumer都会向Coordinator发送join-group，请求重新加入Group（那些原本已经在Group内的也不例外），同时放弃掉已分配给自己的Partition。

SYNC
这一步需要做的事情是：

Coordinator在所有Consumer里选择一个担任Leader，并由Leader调用Partition分配规则来确定消费对应关系。
各个Consumer发送sync-group请求。Leader发送的请求里包含有已经确定的消费分配信息，其他Consumer的请求为空。
Coordinator将消费分配信息原样封装在sync-group响应中，并投递给各个Consumer，最终使Group内所有成员都获知自己该消费的Partition。

在2.4版本特别提出了Incremental（增量）Rebalance

https://www.jianshu.com/p/a8461707d6ea

初始方案
在 Kafka 最初始的解决方案中，是依赖 Zookeeper 的 Watcher 实现的。该方案中，每个 Consumer Group 在 Zookeeper 下都维护了一个对应的 /consumers/{group_id}/ids 路径，该路径下使用临时节点记录该 Consumer Group 中的 Consumer Id，这个 Consumer Id 临时节点在 Consumer 启动时创建。另外，kafka 还会创建 owners 和 offsets 两个节点，这两个节点与 ids 节点同级，其中 owners 记录了 consumer 与 partition 的分配关系；offsets 节点用来记录了对应 Consumer Group 在相应 partition 上的消费位置。

Kafka 在后续版本对 Rebalance 方案进行了改进（也就是 Eager Rebalance Protocol），改进方案的核心设计思想是：将全部的 consumer group 分成多个子集，每个 consumer group 集合在 broker 对应一个 GroupCoordinator，由 GroupCoordinator 管理对应 consumer groups 的 rebalance（每个 broker 都拥有成为 GroupCoordinator 的能力）。

在 kafka 2.4 版本中，为了进一步减少 rebalance 带来的 Stop The World，提出了 Incremental Cooperative Rebalance 协议。其核心思想就是使用将一次全局的 rebalance，改成多次小规模 rebalance，最终收敛到 rebalance 的状态。

 Incremental Cooperative Rebalance 协议，该协议最核心的思想就是：

consumer 比较新旧两个 partition 分配结果，只停止消费回收（revoke）的 partition，对于两次都分配给自己的 partition，consumer 根本没有必要停止消费，这也就解决了 Stop The World 的问题。

通过多轮的局部 rebalance 来最终实现全局的 rebalance 
https://www.bilibili.com/read/cv11642446

Rebalance 流程
每个 ConsumerGroup 有五种状态：

Empty：无任何活跃 Consumer 存在；
Stable：已完成 Rebalance，可供稳定消费；
PreparingRebalance：情况发生变化，有新成员加入或旧成员心跳丢失，需要重新 Balance，要求所有成员重新加入 Group；
CompletingRebalance：所有成员均已入组，各成员等待分配计划；
Dead：Group 生命周期结束，可能因为 session 过期，或者 Group 迁移到其他 Group Coordinator；


Co-partitioning：比如对两个 Topic 做 join，需要将两个 Topic 按相映射的 Partition 来分配给同一个 Consumer；
Sticky Partitioning：对于有状态的 Consumer，希望重启后仍能恢复原有的 Partition 关系而不要 rebalance；
Redundant partitioning：比如搜索引擎需要建多份冗余的索引，希望能使单个 Partition 被多个 Consumer 所消费；
Metadata-based assignment：比如让 Consumer Group 做到 rack aware，消费来自本机架的 Partition；
https://blog.csdn.net/u013474436/article/details/109599966


__consumer_offsets topic
__consumer_offsets 是 Kafka 内部使用的一个 topic，专门用来存储 group 消费的情况，默认情况下有50个 partition，每个 partition 三副本
https://blog.csdn.net/weixin_33970380/article/details/113316660



https://blog.csdn.net/u010022158/article/details/106271208/



