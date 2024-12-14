package main

import "fmt"

func main() {
	//a := 10
	//p := &a
	//
	//fmt.Println(a)
	//fmt.Println(p)
	//fmt.Printf("%v\n", a)
	//fmt.Printf("%v\n", p)
	//fmt.Printf("%p\n", a)
	//fmt.Printf("%p\n", p)



	// compare
	a := 1
	b := 2
	compare(&a, b)
	fmt.Printf("a = %d\n", a)
	a = 3
	compare(&a, b)
	fmt.Printf("a = %d\n", a)
}

func compare(a *int, b int) {
	if *a < b {
		*a = b
	}
}
