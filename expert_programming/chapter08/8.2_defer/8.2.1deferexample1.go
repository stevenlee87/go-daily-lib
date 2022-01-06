package main

import "fmt"

func main() {
	var aInt = 1

	defer fmt.Println(aInt)

	aInt = 2
}

/*
OutPut:
1

延迟函数fmt.Println(aInt)的参数在defer语句出现时就已经确定了，所以无论后面如何修改aInt变量都不会影响延迟函数

*/
