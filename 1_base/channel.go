package main

import "fmt"

func main() {

	// ** 通道 **
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x := <-c // 从通道 c 接收
	y := <-c

	fmt.Println(x, y, x+y)

	// 缓冲区
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 遍历和关闭
	cha := make(chan int, 10)

	go fibonacci(cap(cha), cha)

	for v := range cha {
		fmt.Println(v) // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个数据之后就关闭了通道，所以这里 range 函数在接收到 10 个数据之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候就阻塞了。
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	//close(ch)
}
