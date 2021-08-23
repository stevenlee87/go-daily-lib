package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

/*
解析：
跟上⼀题⼀样，不同的是 *Student 的定义后本身没有初始化值，所以 *Student 是nil的，
但是 *Student 实现了 People 接⼝，接⼝不为 nil 。
*/
