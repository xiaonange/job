package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"job/Kafaka/client"
	"log"
)

const Topic = "testTopic"
const Broker = "192.168.8.200:9092"
const NumPartition = 2
const ReplicationFactor = 1
const ConsumerGroup1 = "consumerTest1"

//集群地址
var address = []string{Broker}

//创建kafka任务发布者
func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	clients, err := sarama.NewClient(address, config)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
	}
	producer, err := sarama.NewSyncProducerFromClient(clients)
	if err != nil {
		log.Fatalf("unable to create kafka producer: %q", err)
	}
	defer producer.Close()
	text := fmt.Sprintf("message %08d", 10)
	partition, offset , err := producer.SendMessage(&sarama.ProducerMessage{Topic: Topic, Key: nil, Value: sarama.StringEncoder(text)})
	if err != nil {
		log.Fatalf("unable to produce message: %q", err)
	}
	fmt.Println(partition,offset)

	//client.KafkaProducer(address)
	client.NewKafkaConsumer(address)
}
