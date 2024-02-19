package main

import "fmt"

/*
原本由 gofmt 格式化后：
var s = "煎鱼进脑子了"

改为 gofumpt 格式化后：
s := "煎鱼进脑子了"
*/
func main() {
	s := "煎鱼进脑子了"
	fmt.Println(s)
}
