package main

import "fmt"

func main() {
	var i = 0

	defer func() {
		fmt.Println(i)
	}()
}

/*
OutPut:
0

延迟函数体中的变量i在编译时绑定原函数中的i，所以延迟函数最终获取的是i的最终值。

*/
