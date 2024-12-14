package main

import (
	"fmt"
	"golang.org/x/exp/mmap"
	"log"
)

func main() {

	// mmap read
	mmapRead()
}

func mmapRead() {
	file, err := mmap.Open("./tmp.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	buff := make([]byte, 2)
	n, err := file.ReadAt(buff, 4)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	fmt.Printf("Read %d bytes: %s\n", n, string(buff))
}

// $ echo "abcdefg" > tmp.txt
