package main

import "fmt"

func main() {
	a := int(100)
	var f *float64

	if f == &a { // invalid operation: f == &a (mismatched types *float64 and *int)
		fmt.Println("f equal &a")
	}
}

/*
Go 的指针多了一些限制。但这也算是 Go 的成功之处：既可以享受指针带来的便利，又避免了指针的危险性。

限制一：Go 的指针不能进行数学运算。
限制二：不同类型的指针不能相互转换。
限制三：不同类型的指针不能使用 == 或 != 比较。
限制四：不同类型的指针变量不能相互赋值。
*/
