package main

import (
	"fmt"
	"github.com/thep0y/go-logger/log"
	"io"
	"net/http"
	"os"
)

func SaveFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err != nil {
		return
	}
	fd, err := os.OpenFile("temp/"+header.Filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer fd.Close()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//buf := make([]byte,128)
	//rw := bufio.NewReadWriter(bufio.NewReader(file),bufio.NewWriter(fd))

	io.Copy(fd, file)
	//for {
	//	count, err := rw.Read(buf)
	//	if err != nil {
	//		break
	//	}
	//	fmt.Println(string(buf[:count]))
	//	rw.Write(buf[:count])
	//}
	//rw.Flush()
}

func main() {
	http.HandleFunc("/test/upload", func(w http.ResponseWriter, r *http.Request) {
		SaveFile(w, r)
	})
	logger := log.NewLogger()
	logger.Info("hhh")
	logger.Error("hhh")
	http.ListenAndServe(":8080", nil)
}
