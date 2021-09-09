package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//conn, err := net.Dial("tcp", "127.0.0.1:8080")
	//defer conn.Close()
	//if err != nil {
	//	log.Panicf("连接服务端失败，err: %v",err)
	//}
	//
	//conn.Write([]byte("hello server!"))
	//
	//buf := make([]byte,1024,1024)
	//count, err := conn.Read(buf)
	//
	//if err != nil {
	//	log.Printf("读取数据失败，err:%v",err)
	//} else {
	//	log.Printf("读取字节数：%d,读取内容：%v",count,string(buf[:count]))
	//}

	go func() {
		fmt.Println(runtime.NumGoroutine())
	}()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
}
