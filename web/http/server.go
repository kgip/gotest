package main

import (
	"fmt"
	"net/http"
)


func main() {
	//http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println(request)
	//	writer.Write([]byte("hello!"))
	//	writer.WriteHeader(http.StatusOK)
	//})
	if res, err := http.Get("https://www.kancloud.cn/aceld/golang/1958304"); err != nil {
		fmt.Println(err)
	} else {
		buf := make([]byte,1024)
		if count, err := res.Body.Read(buf); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(buf[:count]))
		}
	}
	//http.ListenAndServe("localhost:80", nil)
}
