package future

type TaskWithReturn func() interface{}

type ICompletableFuture interface {

}

type result struct {
	index int
	value interface{}
}

type CompletableFuture struct {
	tasks []TaskWithReturn //任务列表
	results []interface{} //任务返回结果集
	resultsChan chan *result //结果集通道
	endChan chan bool //每个任务完成时从该通道拿出一个值，当拿出最后一个值时，关闭workChan
	currFinishedCount int //当前已完成的任务数
}

func NewCompletableFuture(tasks ...TaskWithReturn) *CompletableFuture {
	endChan := make(chan bool,len(tasks))
	for i := 0; i < len(tasks); i++ {
		endChan<- i==len(tasks)-1
	}
	close(endChan)
	return &CompletableFuture{
		tasks: tasks,
		results: make([]interface{},len(tasks)),
		resultsChan: make(chan *result,len(tasks)),
		endChan: endChan,
		currFinishedCount: 0,
	}
}

func (future *CompletableFuture) RunAsync() *CompletableFuture{
	for i := 0; i < len(future.tasks); i++ {
		go future.runTaskAsync(future.tasks[i],i)
	}
	//从通道resultsChan中获取任务返回结果
	for res := range future.resultsChan {
		future.results[res.index] = res.value
	}
	return future
}

//异步运行一个任务
func (future *CompletableFuture) runTaskAsync(task TaskWithReturn,index int) *CompletableFuture  {
	future.resultsChan<- &result{
		index: index,
		value: task(),
	}
	if <-future.endChan {
		close(future.resultsChan)
	}
	return future
}

func (future *CompletableFuture) GetResultByIndex(i int) interface{} {
	return future.results[i]
}

func (future *CompletableFuture) GetAllResults() []interface{} {
	return future.results
}