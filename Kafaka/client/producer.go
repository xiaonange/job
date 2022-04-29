package client

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/*
消息生产者
*/
func KafkaProducer(address []string) {
	//设置配置
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Net.DialTimeout = 10
	//随机的分区类型
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_11_0_0

	//使用配置,新建一个异步生产者
	fmt.Println(1111)
	producer, e := sarama.NewAsyncProducer(address, config)
	fmt.Println(2222)
	if e != nil {
		panic(e)
	}
	defer producer.AsyncClose()

	//发送的消息,主题,key
	msg := &sarama.ProducerMessage{
		Topic: "logstash_test",
		Key:   sarama.StringEncoder("test"),
	}

	var value string
	for {
		value = "this is a message"
		//设置发送的真正内容
		fmt.Scanln(&value)
		//将字符串转化为字节数组
		msg.Value = sarama.ByteEncoder(value)
		fmt.Println(value)

		//使用通道发送
		producer.Input() <- msg

		//循环判断哪个通道发送过来数据.
		select {
		case suc := <-producer.Successes():
			fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
		case fail := <-producer.Errors():
			fmt.Println("err: ", fail.Err)
		}
	}
}
