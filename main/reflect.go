package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	// ** 反射 **

	// 接口与反射
	file, err := os.OpenFile("E:\\Trivia\\test.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader
	r = file

	var w io.Writer
	w = r.(io.Writer)
	write, err := w.Write([]byte("THIS IS A TEST!\n"))
	if err != nil {
		fmt.Println("write file error", err)
		return
	}
	fmt.Println(write)

	// reflect

	// 1. 已知原有类型（强制转换）
	var f float64 = 1.23

	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.ValueOf(f))

	value := reflect.ValueOf(f)
	pointer := reflect.ValueOf(&f)

	realValue := value.Interface().(float64)      // 可以理解为“强制转换”，但是需要注意，转换的时候如果转换的类型不完全符合，则直接panic
	realPointer := pointer.Interface().(*float64) // Golang 对类型要求非常严格，类型一定要完全符合

	testFloat(realValue)
	testFloatPtr(realPointer)

	// 2. 未知原有类型（遍历探测其Field）
	user := User{1, "Jack", 20}
	getFieldAndMethod(user)

	// 设置实际变量的值
	var num float64 = 1.12345

	numPointerValue := reflect.ValueOf(&num)
	newValue := numPointerValue.Elem()

	fmt.Println("type of value:", newValue.Type())
	fmt.Println("settable of value:", newValue.CanSet())

	newValue.SetFloat(3.14) // 重新赋值
	fmt.Println("new value of num:", num)

	// 进行方法调用
	u := User{2, "Tom", 18}
	uValue := reflect.ValueOf(u)

	// 1. 有参数方法
	hasArgsMethod := uValue.MethodByName("ReflectCallFuncHasArgs") // 一定要指定参数为正确的方法名
	args := []reflect.Value{reflect.ValueOf("Tommy")}
	hasArgsMethod.Call(args)

	// 2. 无参数方法
	noArgsMethod := uValue.MethodByName("ReflectCallFuncNoArgs")
	noArgsMethod.Call(make([]reflect.Value, 0)) // 参数传空数组
}

func testFloat(f float64) {
}

func testFloatPtr(fp *float64) {
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Introduce() {
	fmt.Printf("I am %s\n", u.Name)
}

func getFieldAndMethod(i interface{}) {
	typeOf := reflect.TypeOf(i)
	fmt.Println("the type is", typeOf)

	valueOf := reflect.ValueOf(i)
	fmt.Println("the value is", valueOf)

	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		value := valueOf.Field(i).Interface()
		fmt.Printf("%s %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < typeOf.NumMethod(); i++ {
		method := typeOf.Method(i)
		fmt.Printf("%v %s()\n", method.Type, method.Name)
	}
}

func (u User) ReflectCallFuncHasArgs(name string) {
	fmt.Printf("ReflectCallFuncHasArgs: name=%s, originalName=%s\n", name, u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs\n")
}
