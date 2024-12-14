package main

import (
	"context"
	"fmt"
	"time"
)

// 需要往子协程中传递参数，可以使用 context.WithValue()

type option struct {
	internal time.Duration
}

func doTask(ctx context.Context, n string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop", n)
			return
		default:
			fmt.Println("Doing", n)
			// 在子协程中，使用 ctx.Value("options") 获取到传递的值，读取/修改该值
			option := ctx.Value("options").(*option)
			time.Sleep(option.internal * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// context.WithValue() 创建了一个基于 ctx 的子 Context，并携带了值 options
	ctxV := context.WithValue(ctx, "options", &option{1})
	go doTask(ctxV, "Task1")
	go doTask(ctxV, "Task2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
