package service

import (
	"config"
	"fmt"
	"kafka"
	"os"
	"strings"
	"tcpCs"
	"time"
)

var (
	BookQueue  chan string
	KafkaTopic map[string]string
)

func InitService() {
	BookQueue = make(chan string, 10)
	KafkaTopic = make(map[string]string)
	KafkaTopic["checkBook"] = config.Config.KafkaSetting[config.SrvName].CheckServiceTopic
	go KafkaService() //启动一个监听要处理service的管道gorouting
	HandleService()
}

func HandleService() {
	for {
		select {
		case order := <-tcpCs.ServerOrderChan:
			fmt.Println("server handle service: ", order)
			order = strings.TrimSuffix(order, "\n")
			if order == "check book" {
				fmt.Println("handling cheking...")
				go readDirs(config.Config.LocalSetting[config.SrvName].BookStoreDir)
			}
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

func KafkaService() {
	for {
		select {
		case msg := <-BookQueue:
			kafka.WriteMsgToChan(KafkaTopic["checkBook"], msg)
		default:
			time.Sleep(time.Millisecond * 50)
			//kafka.GetFromKafka(KafkaTopic["checkBook"])
		}
	}
}

func readDirs(fileDir string) {
	f, err := os.OpenFile(fileDir, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("openFile err: ", err)
		return
	}
	defer f.Close()
	//fmt.Println("success")
	fileInfos, err := f.Readdir(-1)
	if err != nil {
		fmt.Println("get dir err: ", err)
		return
	}
	for _, item := range fileInfos {
		if !item.IsDir() {
			s := item.Name()
			idx := strings.LastIndex(s, ".")
			if idx <= 0 {
				continue
			}
			bookName := s[:idx]
			fmt.Println(bookName)
			//将书名放入队列chan放入kafka
			BookQueue <- bookName
			fmt.Println("length: ", len(BookQueue))
		} else {
			readDirs(fileDir + "/" + item.Name())
		}
	}
}
