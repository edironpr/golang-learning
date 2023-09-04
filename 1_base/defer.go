package main

import (
	"errors"
	"fmt"
)

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()

	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
}

func main() {
	_, _ = foo(2, 0)
}
