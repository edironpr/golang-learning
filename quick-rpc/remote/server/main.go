package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

func (cal Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}

func main() {
	_ = rpc.Register(new(Cal)) // 发布 Cal 中满足 RPC 注册条件的方法（Cal.Square）
	rpc.HandleHTTP()           // 注册用于处理 RPC 消息的 HTTP Handler

	port := 1234
	log.Printf("Starting server on port %d\n", port)

	// 监听 1234 端口，等待 RPC 请求。
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}
