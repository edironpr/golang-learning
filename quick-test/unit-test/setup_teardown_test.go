package main

import (
	"fmt"
	"os"
	"testing"
)

// setup、teardown

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm Test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm Test2")
}

// 如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。
func TestMain(m *testing.M) {
	// 在这个测试文件中，包含有2个测试用例，Test1 和 Test2。
	// 调用 m.Run() 触发所有测试用例的执行，并使用 os.Exit() 处理返回的状态码，如果不为0，说明有用例失败。
	// 因此可以在调用 m.Run() 前后做一些额外的准备(setup)和回收(teardown)工作。
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
