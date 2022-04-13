package main

import "fmt"

func main() {
	//type A int
	//var a A = 2
	//fmt.Println(a)

	type Wow[T int | string] int

	var a Wow[int] = 123    // 编译正确
	var b Wow[string] = 123 // 编译正确
	fmt.Println(a)
	fmt.Println(b)
	// var c Wow[string] = "hello" // 编译错误，因为"hello"不能赋值给底层类型int

	type pubg[T int | string] string
	var c pubg[string] = "hello"
	fmt.Println(c)
}
