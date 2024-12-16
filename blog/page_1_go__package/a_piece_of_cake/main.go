package main

import (
	"fmt"
	"go-blog/a_piece_of_cake/util"
)

func main() {
	fmt.Println(util.ValidateUsername("john_doe123"))
	fmt.Println(util.ValidateUsername("a"))
	fmt.Println(util.ValidateUsername("123user"))
}