package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// 程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。os.Args变量是一个字符串（string）的切片（slice）。

	// echo
	s, sep := "", ""

	start1 := time.Now()
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end1 := time.Now()
	fmt.Println(end1.Sub(start1).Milliseconds())

	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[:], " "))
	end2 := time.Now()
	fmt.Println(end2.Sub(start2).Milliseconds())

	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
