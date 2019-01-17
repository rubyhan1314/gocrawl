package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler //Sheduler
	WorkerCount int       //worker的数量
	ItemChan chan Item
}

//接口
type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkerChan(chan Request)
	Run()
}

//接口
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//worker公用一个in，out
	//in := make(chan Request)
	out := make(chan ParseResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out) //创建worker
		//createWorker(out, e.Scheduler)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	//参数seeds的request，要分配任务
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("Got %d  item : %v", itemCount, item)
			//itemCount++

			go func() {
				fmt.Printf("item:%+v\n",item)
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

//创建worker
//func createWorker(in chan Request, out chan ParseResult) {
//func createWorker(out chan ParseResult, s Scheduler) {
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	//in := make(chan Request)
	go func() {
		for {
			//s.WorkerReady(in)
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
