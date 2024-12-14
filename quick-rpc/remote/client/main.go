package main

import (
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	// 同步调用
	//syncCall()

	// 异步调用
	asyncCall()
}

func syncCall() {
	// 创建了 HTTP 客户端 client，并且创建了与 localhost:1234 的链接，1234 恰好是 RPC 服务监听的端口
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:1234")

	var result Result
	// 调用远程方法，第1个参数是方法名 Cal.Square，后两个参数与 Cal.Square 的定义的参数相对应。
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}
	log.Printf("%d^2 = %d\n", result.Num, result.Ans)
}

func asyncCall() {
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:1234")

	var result Result

	// client.Go 是异步调用，因此第一次打印 result，result 没有被赋值。而通过调用 <-asyncCall.Done，阻塞当前程序直到 RPC 调用结束，因此第二次打印 result 时，能够看到正确的赋值。
	asyncCall := client.Go("Cal.Square", 12, &result, nil)
	log.Printf("%d^2 = %d\n", result.Num, result.Ans)

	<- asyncCall.Done
	log.Printf("%d^2 = %d\n", result.Num, result.Ans)
}
