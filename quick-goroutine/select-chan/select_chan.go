package main

import (
	"fmt"
	"time"
)

func main() {
	go doTask("Request")

	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(3 * time.Second)
}

var stop = make(chan bool)

func doTask(n string) {
	// 子协程使用 for 循环定时轮询，如果 stop 信道有值，则退出，否则继续轮询
	for {
		select {
		case <- stop:
			fmt.Println("Stop", n)
			return
		default:
			fmt.Println("Doing", n)
			time.Sleep(1 * time.Second)
		}
	}
}