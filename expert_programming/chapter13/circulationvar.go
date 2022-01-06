package main

import "fmt"

func main() {
	var out []*int

	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}

	fmt.Println("Values:", *out[0], *out[1], *out[2])
}

/*
该题目考察循环变量的绑定问题。题目中的out是一个存储整形指针的切片，在循环中，
每次向切片中追加一个i变量的地址。在Go中，循环变量i只分配一次地址，在3次循环中，
i中存储的值分别为0、1、2、3，但是i的地址没有变化。
Output:
Values: 3 3 3
*/
