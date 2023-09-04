package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("connect failed, err=", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		readString, _ := reader.ReadString('\n')
		lineString := strings.Trim(readString, "\r\n")
		if strings.ToUpper(lineString) == "Q" {
			return
		}
		_, err := conn.Write([]byte(lineString))
		if err != nil {
			fmt.Println("write failed, err=", err)
			continue
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read failed, err=", err)
			continue
		}
		fmt.Println(string(buf[:n]))
	}
}
