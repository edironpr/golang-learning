package main

import (
	"fmt"
	"time"
)

func main() {

	// ** 并发 **
	go say("2")
	say("1")

}

func say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}
