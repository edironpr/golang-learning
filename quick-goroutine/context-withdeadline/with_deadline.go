package main

import (
	"context"
	"fmt"
	"time"
)

// 超时退出可以控制子协程的最长执行时间，那 context.WithDeadline() 则可以控制子协程的最迟退出时间
// 某个时刻退出

func doTask(ctx context.Context, n string) {
	// 在子协程中，使用 select 调用 <-ctx.Done() 判断是否需要退出
	for {
		select {
		case <-ctx.Done():
			// 在子协程中，可以通过 ctx.Err() 获取到子协程退出的错误原因
			fmt.Println("Stop", n, ctx.Err())
			return
		default:
			fmt.Println("Doing", n)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	// WithDeadline 用于设置截止时间。在这个例子中，将截止时间设置为1s后，cancel() 函数在 3s 后调用，因此子协程将在调用 cancel() 函数前结束
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	go doTask(ctx, "Task1")
	go doTask(ctx, "Task2")

	time.Sleep(3 * time.Second)
	fmt.Println("Before Cancel")
	cancel()
	time.Sleep(3 * time.Second)
}
