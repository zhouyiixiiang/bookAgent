package tcpCs

import (
	"fmt"
	"net"
)

var (
	ServerListener  net.Listener
	ServerOrderChan chan string
)

func InitServer(addr string, maxOrderNum int) (err error) {
	ServerListener, err = net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("net.Listen(tcp,addr) err: ", err)
		return err
	}
	ServerOrderChan = make(chan string, maxOrderNum)
	go runServer() //在初始化server链接后，把后面的阻塞监听客户端连接请求go出去
	return
}

func runServer() {
	count := 0
	for {
		count++
		fmt.Println("服务器等待客户端建立连接...")
		// 阻塞监听客户端连接请求,如果成功建立连接，则返回用于通信socket的conn
		conn, err := ServerListener.Accept()
		if err != nil {
			fmt.Println("listerner accept err: ", err)
			return
		}
		go serverConService(conn, count)
	}
}

func serverConService(conn net.Conn, count int) {
	defer conn.Close()
	// 获取连接客户端的网络地址
	addr := conn.RemoteAddr()
	fmt.Printf("服务器线程%d与客户端ip：%s连接成功\n", count, addr)
	for {
		// 读取客户端发送的数据
		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("conn.Read(buffer) err: ", err)
			return
		}
		fmt.Println("服务器读到数据：", string(buffer[:n]))
		ServerOrderChan <- string(buffer[:n])
	}
}
