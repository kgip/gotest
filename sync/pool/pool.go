package main

import (
	"fmt"
	"sync"
)

var pool = &sync.Pool{
	New: func() interface{} {
		return 1
	},
}

func main() {
	v := pool.Get()
	fmt.Println(v)
	ch := make(chan int)
	go func() {
		pool.Put(2)
		ch <- pool.Get().(int)
		close(ch)
	}()
	fmt.Println(pool.Get())
	fmt.Println(<-ch)
}
