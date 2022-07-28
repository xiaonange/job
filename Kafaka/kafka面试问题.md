1.kafka数据丢失问题
 Kafka提供了消息发送的ACK机制，这个ACK机制有三个值可以选择。

 当ACK=0的时候，即消息发送到了leader即确认发送成功，此时并不知道其他replica是否已经将消息持久化了没有，这种情况下极有可能出现消息发送了但是丢失的情况。因为如果此时leader节点宕机，
 其他replica会竞选leader，当某一个replica竞选了leader以后，Kafka内部引入了leader epoach机制进行日志截断，此时如果该replica并没有同步到leader接收到这一条消息，那么这条消息就会丢失。

 当ACK=1的时候，即消息发送到了该partition下的ISR集合内的所有replica内。当ISR集合中有多个replica存在，即使此时leader所在的节点宕机，也不会存在消息丢失的情况。因为partition下的leader默认
 是从ISR集合中产生的，而此时ISR集合内的所有replica已经存储了该条消息，所以丢失的可能性几乎为零。

 当ACK=-1的时候，即消息发送到了该partition下的所有replica内。不管leader所在的节点是否宕机，也不管该ISR下的replica是否只有一个，只要该parition下的replica超过一个，那么该消息就不会丢失。
 在日常情况下，我们默认ACK=1，因为ACK=0消息极有可能丢失，ACK=-1消息发送确认时间太长，发送效率太低

 2.消息重复发送问题

   消费端进行去重解决，消费端增加唯一标识。根据唯一标识做去重处理

 3.重复消费问题


 4.kafka消费测问题

5.topic和parttion的差别
topic 是逻辑上的概念，而 partition 是物理上的概念

6.kafka多分组消费
  多分区才是kafka高性能，高吞吐量的良好配置
  kafka一个分组里面的用户只能消费一个分区


 参考文章：https://segmentfault.com/a/1190000041783254