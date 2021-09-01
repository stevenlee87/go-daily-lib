package main

import (
	"fmt"
)

func main() {
	var x string = nil
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}

// golang 中字符串是不能赋值 nil 的，也不能跟 nil ⽐较。
