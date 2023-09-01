package main

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "en" {
		talk = "hello"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	//var peo People = Student{} // 实现方法的接收者是类型指针，类型值不能复制给接口变量
	//think := "bitch"
	//fmt.Println(peo.Speak(think))
}
