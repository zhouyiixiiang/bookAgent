# book_agent
分布式存储管理系统
## 程序架构
支持CS+BS的通讯结构

* basic：常用数据结构
* config：配置项，配置mysql、kafka、TCP通信、本地设置
* kafka：kafka生产者、消费者管理，用于生成消息队列以供多进程使用
* model：使用gorm管理mysql数据库
* service：具体业务处理模块
* tcpCs：管理tcp/ip通讯，即tcp协议下CS架构的server端

### 业务流程
* 启动服务，根据配置文件初始化配置
* 连接mysql数据库
* 初始化kafka生产者，另起go程通过channel监听是否有写往kafka的业务
* 初始化基于tcp通讯的服务器，另起go程监听是否有连接请求，如果有连接请求，则另起go程建立连接，将传给服务器的通信内容写进channel以供service使用
* 初始化kafka消费者，另起go程从kafka消费数据
* 初始化业务服务，使用select监听多个channel，当具体业务的channel被写入指令，则读取指令，并执行对应操作
* 将操作结果写入通知kafka生产者的channel
* kafka生产者监听到对应channel的写入，从channel中读取信息，将信息写入kafka
* kafka消费者将写入的信息进行消费

## nc使用
nc 127.0.0.1 8000	

## kafka使用
```
# 启动zookeeper
/usr/local/bin/zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
# 启动kafka
/usr/local/bin/kafka-server-start /usr/local/etc/kafka/server.properties
# 启动消费者
/usr/local/bin/kafka-console-consumer --topic=checkService --bootstrap-server=127.0.0.1:9092 --from-beginning
```