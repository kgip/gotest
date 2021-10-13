package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	ctx := context.Background()
	vctx := context.WithValue(ctx, "name", "aaaa")
	wg.Add(1)
	go func(pctx context.Context) {
		fmt.Println("name:", pctx.Value("name"))
		deadline, ok := pctx.Deadline()
		if ok {
			fmt.Println("deadline:", deadline)
		} else {
			fmt.Println("获取deadline失败")
		}
		dvctx, cancel := context.WithDeadline(pctx, time.Now().Add(3*time.Second))
		go func(spctx context.Context) {
			for {
				select {
				case <-spctx.Done():
					fmt.Println("子协程退出了:")
					wg.Done()
					return
				default:
					fmt.Println("子协程执行中")
					time.Sleep(200 * time.Microsecond)
				}
			}
		}(dvctx)
		time.Sleep(time.Second)
		cancel()
	}(vctx)
	wg.Wait()
}
