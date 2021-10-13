package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

type Error struct {
	code    int
	message string
	path    string
}

func (this *Error) Error() string {
	return fmt.Sprintf("错误码：%d,错误信息：%s,错误地址：%s", this.code, this.message, this.path)
}

func ReadFile(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalln("打开文件出错，err=", err)
	}

	buf := make([]byte, 1024)
	count, err := file.Read(buf)
	if err != nil {
		log.Fatalln("读取文件出错，err=", err)
	}
	fmt.Printf("读取文件成功，读取总字节数：%d,读取的文件内容：\n%s", count, string(buf[:count]))
}

func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer func() {
		file.Close()
		err := recover()
		if err != nil {
			log.Println("执行出错，err=", err)
		}
		log.Println("读写文件后")
	}()
	if err != nil {
		log.Panic("打开文件失败，err=", err)
	}
	file.Write([]byte(content))
	err = &Error{
		code:    402,
		message: "服务器错误",
		path:    "/http/server.go",
	}
	panic(err)
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("执行出错，err=", err)
			debug.PrintStack()
		}
	}()
	//ReadFile("http/server.go")
	//WriteFile("io/test.txt","hello!")
	//log.Println("读写文件后")
	//os.Stdout.Write([]byte("hello~\n"))
	buf := make([]byte, 1)
	if count, err := os.Stdin.Read(buf); err == nil {
		fmt.Printf("读取字节数：%d,读取内容：\n%s", count, string(buf[:count]))
	} else {
		fmt.Println("err=", err)
	}
	//srcFile, err := os.Open("defer/defer.go")
	//defer srcFile.Close()
	//if err != nil {
	//	fmt.Println("打开文件出错！ err=",err)
	//	return
	//}
	//dstFile, err := os.OpenFile("dst.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
	//defer dstFile.Close()
	//if err != nil {
	//	fmt.Println("打开文件出错！ err=",err)
	//	return
	//}
	//count, err := io.Copy(dstFile, srcFile)
	//if err!=nil {
	//	if err != nil {
	//		fmt.Println("打开文件出错！ err=",err)
	//		return
	//	}
	//}
	//fmt.Println("count=",count)
	//fd, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY, os.ModePerm)
	//defer fd.Close()
	//if err != nil && os.IsNotExist(err) {
	//	log.Println(err)
	//	return
	//}
	//w := bufio.NewWriter(fd)
	//w.Write([]byte("hfeiahfhifeiawhfiewfheiihfeihwaiofehaiofihfeiwhfiee"))
	//w.Flush()
}
