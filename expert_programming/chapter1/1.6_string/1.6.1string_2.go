package main

import "fmt"

func main() {
	var a string

	// cannot use nil as type string in assignment
	// a = nil // 不存在值为nil的string

	a = "" // 空字符只是长度为0，但不是nil
	fmt.Println(len(a))
}
