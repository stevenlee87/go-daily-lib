package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	fmt.Println(cap(s))
}

/*
解析：
输出为 [0 0 0 0 0 1 2 3]
make 在初始化切⽚时指定了⻓度，所以追加数据时会从 len(s) 位置开始填充数据。
*/
