package main

import (
	"log"
	"net"
)

func main() {
	//创建服务器，对8080端口进行监听
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	log.Println("服务端创建成功，开始监听8080端口")
	if err != nil {
		log.Panicf("服务器创建失败，err:%v",err)
	}

	for  {
		//接收客户端连接
		log.Println("开始接收客户端连接...")
		conn, err := listen.Accept()
		log.Println("接收到一个客户端连接...")
		if err != nil {
			log.Printf("客户端监听失败，err:%v",err)
		} else {
			go func(conn net.Conn) {
				defer conn.Close()
				buf := make([]byte, 1024, 1024)
				count, e := conn.Read(buf)

				if e != nil {
					log.Printf("读取数据失败，err:%v",e)
					conn.Write([]byte("读取数据失败！"))
				} else {
					conn.Write([]byte("hello client!"))
					log.Printf("读取总字节数：%d,读取内容：\n%s",count,string(buf[:count]))
				}
			}(conn)
		}
	}
}
