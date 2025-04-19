package main

import (
	"fmt"
	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")

	// 链接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.10.7:7344"}, config)
	if err != nil {
		panic("producer closed" + err.Error())
	}
	defer client.Close()
	fmt.Println(11111)

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		panic("send message fail" + err.Error())
	}
	fmt.Println(pid, offset)
}
