package main

import (
	"context"
	"fmt"
	"time"
)

// 如果需要控制子协程的执行时间，可以使用 context.WithTimeout 创建具有超时通知机制的 Context 对象
// 持续多久退出

func doTask(ctx context.Context, n string) {
	// 在子协程中，使用 select 调用 <-ctx.Done() 判断是否需要退出
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop", n)
			return
		default:
			fmt.Println("Doing", n)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// WithTimeout()的使用与 WithCancel() 类似，多了一个参数，用于设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go doTask(ctx, "Task1")
	go doTask(ctx, "Task2")

	time.Sleep(3 * time.Second)
	fmt.Println("Before Cancel")
	cancel()
	time.Sleep(3 * time.Second)
}
