package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var (
	ClientProducer sarama.SyncProducer
	KafkaMsgChan   chan *KafkaMsg
)

type KafkaMsg struct {
	topic string
	data  string
}

func InitProducer(addrs []string, kafkamaximumSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 连接kafka
	ClientProducer, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("sarama.NewSyncProducer err: ", err)
		return err
	}
	//defer  clientKafka.Close()
	KafkaMsgChan = make(chan *KafkaMsg, kafkamaximumSize)
	go SendToKafka()
	return
}

func WriteMsgToChan(topic, data string) {
	msg := &KafkaMsg{
		topic: topic,
		data:  data,
	}
	KafkaMsgChan <- msg
}

func SendToKafka() (err error) {
	for {
		select {
		case mg := <-KafkaMsgChan:
			// 构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = mg.topic
			msg.Value = sarama.StringEncoder(mg.data)
			// 发送消息到kafka
			pid, offset, err := ClientProducer.SendMessage(msg)
			if err != nil {
				return err
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
