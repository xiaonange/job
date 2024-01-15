1.kafka数据丢失问题
 Kafka提供了消息发送的ACK机制，这个ACK机制有三个值可以选择。

 当ACK=0的时候，即消息发送到了leader即确认发送成功，此时并不知道其他replica是否已经将消息持久化了没有，这种情况下极有可能出现消息发送了但是丢失的情况。因为如果此时leader节点宕机，
 其他replica会竞选leader，当某一个replica竞选了leader以后，Kafka内部引入了leader epoach机制进行日志截断，此时如果该replica并没有同步到leader接收到这一条消息，那么这条消息就会丢失。

 当ACK=1的时候，即消息发送到了该partition下的ISR集合内的所有replica内。当ISR集合中有多个replica存在，即使此时leader所在的节点宕机，也不会存在消息丢失的情况。因为partition下的leader默认
 是从ISR集合中产生的，而此时ISR集合内的所有replica已经存储了该条消息，所以丢失的可能性几乎为零。

 当ACK=-1的时候，即消息发送到了该partition下的所有replica内。不管leader所在的节点是否宕机，也不管该ISR下的replica是否只有一个，只要该parition下的replica超过一个，那么该消息就不会丢失。
 在日常情况下，我们默认ACK=1，因为ACK=0消息极有可能丢失，ACK=-1消息发送确认时间太长，发送效率太低

 2.消息重复发送问题

   生产端代码逻辑做检查，消费端进行去重解决，消费端增加唯一标识。根据唯一标识做去重处理

 3.重复消费问题
 消费端进行去重解决，消费端增加唯一标识。根据唯一标识做去重处理

 4.kafka消费测问题

5.topic和parttion的差别
topic 是逻辑上的概念，而 partition 是物理上的概念

6.kafka多分组消费
  多分区才是kafka高性能，高吞吐量的良好配置
  kafka一个分组里面的用户只能消费一个分区

7.kafka的rebalance机制
rebalance就是说如果消费组里的消费者数量有变化或消费的分区数有变化，kafka会重新分配消费者消费分区的关系。

如下情况可能会触发消费者rebalance
1.消费组里的consumer增加或减少了
2.动态给topic增加了分区
3.消费组订阅了更多的topic

要有三种rebalance的策略：range()、round-robin(轮询)、sticky(粘性)。



8.kafka概念
Producer：消息生产者，向 Kafka Broker 发消息的客户端。

Consumer：消息消费者，从 Kafka Broker 取消息的客户端。

Consumer Group：消费者组（CG），消费者组内每个消费者负责消费不同分区的数据，提高消费能力。一个分区只能由组内一个消费者消费，消费者组之间互不影响。所有的消费者都属于某个消费者组，即消费者组是逻辑上的一个订阅者。

Broker：一台 Kafka 机器就是一个 Broker。一个集群由多个 Broker 组成。一个 Broker 可以容纳多个 Topic。

Topic：可以理解为一个队列，Topic 将消息分类，生产者和消费者面向的是同一个 Topic。

Partition：为了实现扩展性，提高并发能力，一个非常大的 Topic 可以分布到多个 Broker （即服务器）上，一个 Topic 可以分为多个 Partition，每个 Partition 是一个 有序的队列。

Replica：副本，为实现备份的功能，保证集群中的某个节点发生故障时，该节点上的 Partition 数据不丢失，且 Kafka 仍然能够继续工作，Kafka 提供了副本机制，一个 Topic 的每个分区都有若干个副本，一个 Leader 和若干个 Follower。

Leader：每个分区多个副本的“主”副本，生产者发送数据的对象，以及消费者消费数据的对象，都是 Leader。

Follower：每个分区多个副本的“从”副本，实时从 Leader 中同步数据，保持和 Leader 数据的同步。Leader 发生故障时，某个 Follower 还会成为新的 Leader。

Offset：消费者消费的位置信息，监控数据消费到什么位置，当消费者挂掉再重新恢复的时候，可以从消费位置继续消费。

ZooKeeper：Kafka 集群能够正常工作，需要依赖于 ZooKeeper，ZooKeeper 帮助 Kafka 存储和管理集群信息

10.kafka副本同步机制
  1.leader读写请求:kafka是由follower周期性或者尝试去pull(拉)过来
  2.leader有isr（inser）列表，如果follower有故障，则放到osr中


 参考文章：https://segmentfault.com/a/1190000041783254