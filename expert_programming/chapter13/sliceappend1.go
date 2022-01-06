package main

import "fmt"

func main() {
	x := make([]int, 0, 10) // 长度为0，空间为10的切片x
	fmt.Printf("x1 len is:%d, cap is:%d\n", len(x), cap(x))
	x = append(x, 1, 2, 3)
	fmt.Printf("x2 len is:%d, cap is:%d\n", len(x), cap(x))
	y := append(x, 4)
	fmt.Printf("y1 len is:%d, cap is:%d\n", len(y), cap(y))
	z := append(x, 5)
	fmt.Printf("x2 len is:%d, cap is:%d\n", len(x), cap(x))
	fmt.Printf("z1 len is:%d, cap is:%d\n", len(z), cap(z))

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

/*
annotation：
当append向切片x追加元素时，在空间足够存放新元素的情况下，新元素将从x[len(x)]位置开始存放，append会
生成一个新的切片，但不会修改原切片x。
Output:
[1 2 3]
[1 2 3 5]
[1 2 3 5]
*/
