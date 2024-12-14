package main

import "log"

type Result struct {
	Num, Ans int
}

type Cal int

func (cal Cal) Square(num int) *Result {
	return &Result{
		Num: num,
		Ans: num * num,
	}
}

func main() {
	// Local
	c := new(Cal)
	result := c.Square(12)
	log.Printf("%d^2 = %d\n", result.Num, result.Ans)
}
