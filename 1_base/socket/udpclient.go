package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 8002})
	if err != nil {
		fmt.Println("connect failed, err=", err)
		return
	}

	defer socket.Close()

	_, err = socket.Write([]byte("hello server"))
	if err != nil {
		fmt.Println("write failed, err=", err)
	}

	data := make([]byte, 4096)
	n, addr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read failed, err=", err)
		return
	}
	fmt.Printf("data: %v addr: %v\n", string(data[:n]), addr)
}
