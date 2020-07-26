package main

import (
	"config"
	_ "github.com/go-sql-driver/mysql"
	"kafka"
	"model"
	"service"
	"tcpCs"
)

func main() {
	config.Init("/Users/zhouyixiang/Documents/workspace/src/book_agent/config.json")
	model.Init()
	kafka.InitProducer(config.Config.KafkaSetting[config.SrvName].Addrs, config.Config.KafkaSetting[config.SrvName].MaximumChanSize)
	tcpCs.InitServer(config.Config.TCPSetting[config.SrvName].ServerAddr, config.Config.TCPSetting[config.SrvName].ServerMaxOrderChanNum)
	kafka.InitConsumer(config.Config.KafkaSetting[config.SrvName].Addrs, config.Config.KafkaSetting[config.SrvName].CheckServiceTopic)
	go kafka.GetFromKafka(config.Config.KafkaSetting[config.SrvName].CheckServiceTopic)
	service.InitService()
}
