package main

import "fmt"

// 下⾯代码输出什么？
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/*
defer 在定义的时候会计算好调⽤函数的参数，所以会优先输出 10 、 20 两个参
数。然后根据定义的顺序倒序执⾏。
*/
