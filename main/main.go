package main

import (
	"fmt"
	"unsafe"
)

var (
	x int8
	y bool
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(x, y)

	var a = 1

	fmt.Println(a, true)

	m := "hello"
	n := 1
	fmt.Println(m, n)

	const name = "Jack"
	const open bool = false
	fmt.Println(name, open)

	const (
		j = "abc"
		k = len(j)
		l = unsafe.Sizeof(j)
	)
	fmt.Println(j, k, l)

	const (
		q = iota
		p
		u
	)
	fmt.Println(q, p, u)

	const (
		a1 = iota //0
		a2        //1
		a3        //2
		a4 = "ha" //独立值，iota += 1
		a5        //"ha"   iota += 1
		a6 = 100  //iota +=1
		a7        //100  iota +=1
		a8 = iota //7,恢复计数
		a9        //8
	)
	fmt.Println(a1, a2, a3, a4, a5, a6, a7, a8, a9)

	const (
		i1 = 1 << iota // 1 * 2^0
		i2 = 3 << iota // 3 * 2^1
		i3             // 3 * 2^2
		i4             // 3 * 2^3
	)

	fmt.Println("i1=", i1)
	fmt.Println("i2=", i2)
	fmt.Println("i3=", i3)
	fmt.Println("i4=", i4)

	nums := []int{1, 3, 5, 7}
	for i, num := range nums {
		fmt.Printf("%d -> %d\n", i, num)
	}

	kvs := map[string]string{"a": "apple", "b": "bear"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	var m1 int = 100
	var m2 int = 200
	var result int
	result = max(m1, m2)

	fmt.Println(result)

	fmt.Println(swap("Google", "Bing"))

	var c Circle
	c.radius = 10
	fmt.Println("the area of circle is ", c.getArea())

	DeferFunction()

	arrAssign(10)
	// 产生错误后 程序继续
	fmt.Println("continue...")

	// 数组
	balance := [10]float64{1.0, 2.0, 3.0, 4.0}
	fmt.Println(balance)

	salary := [...]float32{1.1, 2.2, 3.3}
	fmt.Println(salary)

	mark := []int{1, 2, 3, 4}
	fmt.Println(mark)

	names := [5]string{1: "Jack", 3: "Alice"}
	fmt.Println(names)
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

// Circle /* 定义结构体 */
type Circle struct {
	radius float64
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func DeferFunction() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

func arrAssign(i int) {
	var arr [10]int

	// 错误拦截要在产生错误前设置
	defer func() {
		// 设置 recover 拦截错误信息
		err := recover()

		// 产生 panic 异常  打印错误信息
		if err != nil {
			fmt.Println(err)
		}
	}()

	arr[i] = 10
}
