package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doTask(i + 1)
	}

	// wg.Wait() 会等待所有的子协程任务全部完成，所有子协程结束后，才会执行 wg.Wait() 后面的代码
	wg.Wait()
	fmt.Println("All Task Done")
}

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}