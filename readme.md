# book_agent
图书管理系统
## 程序架构
支持CS+BS的通讯结构

* server端接收client传来的指令
* server按照client传输的指令作出相应操作
	* check book 检查本地已有图书
		* 遍历指定文件夹内的书名，将遍历结果放入kafka
	* 

## nc使用
nc 127.0.0.1 8000	

## kafka使用
```
# 消费者
```