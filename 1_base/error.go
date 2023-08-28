package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("square result can not be negative")
	}

	return f, nil
}
