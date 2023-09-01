package main

import "fmt"

func main() {
	a := 10
	p := &a

	fmt.Println(a)
	fmt.Println(p)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", p)
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", p)
}
