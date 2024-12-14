package main

import (
	"context"
	"fmt"
	"time"
)

// Go 语言提供了 Context 标准库可以解决这类场景的问题，Context 的作用和它的名字很像，上下文，即子协程的下上文。Context 有两个主要的功能：
//	1.通知子协程退出（正常退出，超时退出等）
//	2.传递必要的参数

// 使用 Context 改写，效果与 select+chan 相同
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
	// 控制单个协程

	// context.Background() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context
	ctxBg := context.Background()
	// context.WithCancel(parent) 创建可取消的子 Context，同时返回函数 cancel
	ctx, cancel := context.WithCancel(ctxBg)

	go doTask(ctx, "Connect")

	time.Sleep(3 * time.Second)

	// 主协程中，调用 cancel() 函数通知子协程退出
	cancel()

	time.Sleep(3 * time.Second)


	// 控制多个协程

	ctxMulti, cancelMulti := context.WithCancel(context.Background())
	go doTask(ctxMulti, "Task1")
	go doTask(ctxMulti, "Task2")
	time.Sleep(3 * time.Second)
	cancelMulti()
	time.Sleep(3 * time.Second)
}
