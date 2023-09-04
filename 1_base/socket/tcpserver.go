package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err=", err)
			break
		}
		s := string(buf[:n])
		fmt.Println("receive from client, data=", s)
		conn.Write([]byte(s))
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("listen failed, err=", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err=", err)
			continue
		}
		go process(conn)
	}
}
