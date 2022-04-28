package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return talk
}

//接口实现
func main() {
	var stu  Student
	think := "bitch"
	fmt.Println(stu.Speak(think))
}
