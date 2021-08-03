package main

import "fmt"

func main() {
	s := "Hello World!"
	fmt.Println(s[0])

	fmt.Println(s[:5])
}

// 72
// Hello
// 上篇文章你不知道的 Go 之 slice中也介绍过了，字符串的切片操作返回的不是切片，而是字符串。
