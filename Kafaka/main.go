package main

import "job/Kafaka/client"

const Topic = "testTopic"
const Broker = "120.25.121.92:9002"
const NumPartition = 2
const ReplicationFactor = 1
const ConsumerGroup1 = "consumerTest1"

//集群地址
var address = []string{Broker}

//创建kafka任务发布者
func main() {

	client.KafkaProducer(address)
	client.NewKafkaConsumer(address)
}
