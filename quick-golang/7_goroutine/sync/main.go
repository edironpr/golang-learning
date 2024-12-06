package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()               // count -1
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1) // count +1
		go download("a.com/" + strconv.Itoa(i))
	}
	wg.Wait() // 阻塞，直到 count = 0
	fmt.Println("Done!")
}
