package main

import "fmt"

type Phone interface {
	call()
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("This is iPhone")
}

type Samsung struct {
}

func (samsung Samsung) call() {
	fmt.Println("This is Samsung")
}

func main() {

	// ** 接口 **
	var phone Phone

	phone = IPhone{}
	phone.call()

	phone = Samsung{}
	phone.call()

	// 类型断言
	var i interface{} = "a"
	fmt.Println(i.(string))

	value, ok := i.(string)
	if ok {
		fmt.Println("The value is", value)
	} else {
		fmt.Println("It's not ok for type string")
	}

	switch t := i.(type) {
	default:
		fmt.Printf("unexpected type %T", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("int %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to int %d\n", *t)
	}

	if i == nil {
		fmt.Printf("NULL\n")
	} else if value, ok := i.(bool); ok {
		fmt.Printf("boolean %t\n", value)
	} else if value, ok := i.(int); ok {
		fmt.Printf("int %d\n", value)
	} else if value, ok := i.(string); ok {
		fmt.Printf("string %s\n", value)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v\n", i, i))
	}

}
