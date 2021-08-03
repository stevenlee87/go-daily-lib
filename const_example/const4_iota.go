package main

import "fmt"

const (
	A int = 1
	B int = 2
	C int = iota + 1
	D
	E
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
}

/*
上面iota出现在第 2 行（从 0 开始），C的值为2 + 1 = 3。D和E分别为 4, 5。

一定要注意iota的值等于它出现在组中的第几行，而非它的第几次出现。

输出：
1
2
3
4
5
*/
