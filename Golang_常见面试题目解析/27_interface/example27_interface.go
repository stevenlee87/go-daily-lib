package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "foolish" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "foolish"
	fmt.Println(peo.Speak(think))
}

/*
cannot use Student{} (type Student) as type People in assignment:
	Student does not implement People (Speak method has pointer receiver)

编译失败，值类型 Student{} 未实现接⼝ People 的⽅法，不能定义为 People 类型。

在 golang 语⾔中， Student 和 *Student 是两种类型，第⼀个是表示 Student 本身，
第⼆个是指向 Student 的指针。
*/
