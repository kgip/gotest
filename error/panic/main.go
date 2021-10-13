package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/thep0y/go-logger/log"
)

func test() {
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
			if err, ok := e.(error); ok {
				log.Error(err.Error())
			} else if err, ok := e.(string); ok {
				log.Error(err)
			}
		}
	}()
	panic("error with string !!!")
}

func GetMD5(content string) string {
	MD5 := md5.New()
	MD5.Write([]byte(content))
	return hex.EncodeToString(MD5.Sum(nil))
}

func main() {
	test()
	fmt.Println(GetMD5("afegg"))
}
