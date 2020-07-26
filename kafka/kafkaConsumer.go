package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	ClientConsumer sarama.Consumer
	KafkaOutChan   chan string
)

func InitConsumer(addrs []string, topic string) (err error) {
	// 连接kafka
	ClientConsumer, err = sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Println("sarama.NewConsumer err: ", err)
		return err
	}
	//defer  clientKafka.Close()
	return
}

func GetFromKafka(topic string) {
	// 根据topic取到所有的分区
	partitionList, err := ClientConsumer.Partitions(topic)
	if err != nil {
		fmt.Println("ClientConsumer.Partitions(topic) err: ", err)
		return
	}
	for partition := range partitionList {
		// 针对每一个分区创建一个对应的分区消费者
		pc, err := ClientConsumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("ClientConsumer.ConsumePartition err: ", err)
			return
		}
		//异步从每个分区消费信息
		go ConsumeMessage(pc)
	}

}

func ConsumeMessage(pc sarama.PartitionConsumer) {
	defer pc.AsyncClose()
	for msg := range pc.Messages() {
		fmt.Printf("partion:%d offset:%d key:%v value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}
