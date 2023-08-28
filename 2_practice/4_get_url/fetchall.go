package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.3f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("when get %s: %v", url, err) // send to channel ch
		return
	}

	bytesCount, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("when read %s: %v", url, err)
		return
	}

	duration := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.3fs\t%d\t%s", duration, bytesCount, url)
}
