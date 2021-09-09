package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	rwLock sync.RWMutex
	lock sync.Mutex
	once = &sync.Once{}
)

type Student struct {
	Name string
	Age int
}

func main() {
	//ticker := ticker.NewTicker(time.Second, func() {
	//	fmt.Println(time.Now())
	//}).Do()
	//
	//time.Sleep(10*time.Second)
	//
	//ticker.Stop()
	//fmt.Println("定时器停止执行...")
	//wg.Add(3)
	//go func() {
	//	time.Sleep(time.Second)
	//	fmt.Println("任务1执行")
	//	wg.Done()
	//}()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("任务2执行")
	//	wg.Done()
	//}()
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("任务3执行")
	//	wg.Done()
	//}()
	//wg.Wait()
	//var state int64 = 0
	////fmt.Printf("%p\n",&state)
	//workerCount := 1000
	//wg.Add(workerCount)
	//for i := 0; i < workerCount; i++ {
	//	go func() {
	//		defer wg.Done()
	//		//fmt.Printf("%p",&state)
	//		//state = state+1
	//		for !atomic.CompareAndSwapInt64(&state, state, state+1) {
	//			fmt.Println("增加操作执行失败！")
	//		}
	//	}()
	//}
	//wg.Wait()
	//fmt.Println("任务执行完成",state)
	//noBufChan := make(chan int)
	////wg.Add(1)
	//go func() {
	//	//defer wg.Done()
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("管道值被写入")
	//		noBufChan<-1
	//	}
	//	close(noBufChan)
	//}()
	////fmt.Printf("管道值被读取：%v",<-noBufChan)
	//
	//for {
	//	select {
	//		case v := <-noBufChan:
	//			if v == 0 {
	//				return
	//			}
	//			fmt.Println("管道值被读取:",v)
	//		default:
	//			fmt.Println("管道为空")
	//	}
	//}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("函数执行")
			})
			wg.Done()
		}()
	}
	wg.Wait()



	//wg.Wait()
}
