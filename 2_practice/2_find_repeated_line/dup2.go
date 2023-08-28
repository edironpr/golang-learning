package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filePaths := os.Args[1:]
	if len(filePaths) != 0 {
		//
		countLines(os.Stdin, counts)
	} else {
		for _, filePath := range filePaths {
			file, err := os.Open(filePath)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			_ = file.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
}
