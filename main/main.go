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

	// ** 指针 **

	// 取址符
	var aa int = 1
	println("the address of variable 'a' is", &aa)

	// 指针类型
	bb := 1.2
	var aaP *int = &aa
	var bbP *float64 = &bb
	println("the address 'aaP' is", aaP)
	println("the address 'bbP' is", bbP)

	// 指针所指变量的值
	println("the value of aaP point is", *aaP)
	println("the value of bbP point is", *bbP)

	// 空指针
	var nptr *int
	println(nptr)
	println(nptr == nil)
	println(nptr != nil)

	// 指针数组
	const arrSize = 3
	arr := [arrSize]int{1, 10, 100}
	var arrPtr [arrSize]*int

	for i := 0; i < arrSize; i++ {
		arrPtr[i] = &arr[i]
		fmt.Printf("arr[%d] = %d\n", i, *arrPtr[i])
	}

	// 指针的指针
	var z int
	var zPtr *int
	var zPptr **int

	z = 18
	zPtr = &z
	zPptr = &zPtr

	println(z)
	println(*zPtr)
	println(**zPptr)

	// 指针参数
	s1 := 100
	s2 := 200
	fmt.Printf("s1 = %d, s2 = %d\n", s1, s2)

	swap2(&s1, &s2)

	fmt.Printf("s1 = %d, s2 = %d\n", s1, s2)

	// ** 结构体 **

	// 定义
	fmt.Println(Book{"A", "jack", "Go", 1})
	fmt.Println(Book{title: "B", author: "tom", subject: "Go", id: 2})
	fmt.Println(Book{title: "C", author: "alice"})

	// 访问成员
	var book1 Book
	book1.title = "Go Tutorial"
	book1.author = "ed"
	book1.subject = "GO"
	book1.id = 3
	fmt.Println(book1)

	book2 := Book{"Go Practice", "ed", "GO", 4}
	fmt.Println(book2.title, book2.author, book2.subject, book2.id)

	// 函数参数
	printBook(Book{"L&P", "ed", "drama", 5})

	// 结构体指针
	var bookPtr *Book
	bookPtr = &book1
	println("the title of the book1 is", bookPtr.title)
	printBook2(&book2)

	// ** 切片 **

	// 定义
	var slice1 []int
	var slice2 = make([]int, 10)
	slice3 := make([]string, 5)
	fmt.Println(slice1)
	printSlice(slice1)
	fmt.Println(slice2)
	printSlice(slice2)
	fmt.Println(slice3)

	// 初始化
	slice4 := arr[:]
	fmt.Println(slice4)
	printSlice(slice4)
	slice5 := arr[0:2]
	fmt.Println(slice5)
	slice6 := arr[1:]
	fmt.Println(slice6)
	slice7 := arr[:1]
	fmt.Println(slice7)
	slice8 := make([]int, 2, 5)
	fmt.Println(slice8)

	// len() cap() 函数
	sli1 := make([]int, 3, 5)
	printSlice(sli1)

	// 空切片
	var sli2 []int
	if sli2 == nil {
		println("sil2 is nil")
	}
	printSlice(sli2)

	// 切片截取
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)
	fmt.Println(numbers)
	fmt.Println(numbers[1:4])
	fmt.Println(numbers[:3])
	fmt.Println(numbers[4:])
	printSlice(numbers[:2])
	printSlice(numbers[2:5])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	// append() copy() 函数
	numbers = append(numbers, 0)
	printSlice(numbers)
	numbers = append(numbers, 1)
	printSlice(numbers)
	numbers = append(numbers, 1, 2, 3)
	printSlice(numbers)

	copyNumbers := make([]int, len(numbers), cap(numbers)*2)
	printSlice(copyNumbers)
	copyNumbers1 := make([]int, 5, cap(numbers)*2)
	copy(copyNumbers1, numbers)
	printSlice(copyNumbers1)

	copy(copyNumbers, numbers)
	printSlice(copyNumbers)

	// ** Map 集合 **

	// 定义
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["CN"] = "China"
	map1["AM"] = "America"
	map1["HK"] = "HongKong"
	map1["JP"] = "Japan"

	for k, v := range map1 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	v1, exist1 := map1["CN"]
	println(v1, exist1)
	v2, exist2 := map1["TW"]
	println(v2, exist2)

	// delete 函数
	delete(map1, "HK")

	for k, v := range map1 {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// ** 类型转换 **
	var count int = 4
	var sum int = 10
	var devideRes float64
	devideRes = float64(sum) / float64(count)
	fmt.Printf("%f\n", devideRes)

	var bo Book
	fmt.Println(bo)
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

func swap2(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

type Book struct {
	title   string
	author  string
	subject string
	id      int
}

func printBook(book Book) {
	fmt.Println("the title of the book is", book.title)
	fmt.Println("the author of the book is", book.author)
	fmt.Println("the subject of the book is", book.subject)
	fmt.Println("the id of the book is", book.id)
}

func printBook2(bookPtr *Book) {
	fmt.Println("the title of the book is", bookPtr.title)
	fmt.Println("the author of the book is", bookPtr.author)
	fmt.Println("the subject of the book is", bookPtr.subject)
	fmt.Println("the id of the book is", bookPtr.id)
}

func printSlice(slice []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)
}
