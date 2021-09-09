package main

import (
	"fmt"
	"strconv"
)

func writeNum(n int,numChan chan int) {
	for i:=1;i<=n;i++ {
		numChan<-i
	}
	close(numChan)
}

func readNum(numChan chan int,resChan chan map[int]int,endChan chan bool)  {
	for i := 0; i < 25; i++ {
		n := <-numChan
		res := make(map[int]int,1)
		res[n] = ((1+n)*n)>>1
		resChan<-res
	}
	if <-endChan {
		close(resChan)
	}
}

var (
	cache = make(map[string]int,10)
	semaphore = make(chan bool,10)
)

func init() {
	for i := 0; i < cap(semaphore); i++ {
		semaphore<- false
	}
}

func addNum(n int) {
	<-semaphore
	cache[strconv.Itoa(n)] = (n+1)*n>>1
	fmt.Println("添加数字：",n)
	semaphore<-false
}



func main() {
	//numChan := make(chan int,200)
	//resChan := make(chan map[int]int,200)
	//endChan := make(chan bool,8)
	//for i := 0; i < 8; i++ {
	//	endChan<- i==7
	//}
	//close(endChan)
	//go writeNum(200,numChan)
	//for i := 0; i < 8; i++ {
	//	go readNum(numChan,resChan,endChan)
	//}
	//for v := range resChan {
	//	for k, i := range v {
	//		fmt.Printf("res[%d]=%d\n",k,i)
	//	}
	//}

	//for i := 0; i < 100; i++ {
	//	go addNum(i)
	//}
	//time.Sleep(4*time.Second)
	//var i interface{} = 1

	//results := future.NewCompletableFuture(func() interface{} {
	//	fmt.Println("任务1执行...")
	//	time.Sleep(time.Second)
	//	return 3
	//}, func() interface{} {
	//	fmt.Println("任务2执行...")
	//	time.Sleep(2*time.Second)
	//	return 4
	//}, func() interface{} {
	//	fmt.Println("任务3执行...")
	//	time.Sleep(3*time.Second)
	//	return 3
	//}).RunAsync().GetAllResults()
	//fmt.Println(results)

	for i := 0; i < 1000; i++ {
		go addNum(i+10)
	}

}
