package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/*
*

	练习 1.10： 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。
*/
func main() {

	url := os.Args[1]
	result := make([]string, 2)

	for i := 0; i < 2; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "when get %s: %v", url, err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "when read %s: %v", url, err)
		}
		result[i] = string(bytes)
		fmt.Printf("%.3f\n", time.Since(start).Seconds())
	}

	if result[0] == result[1] {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}

}
