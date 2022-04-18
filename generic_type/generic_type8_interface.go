package main

import "fmt"

// 空接口代表所有类型的集合。写入类型约束意味着所有类型都可拿来做类型实参
type Slice[T interface{}] []T

var s1 Slice[int]               // 正确
var s2 Slice[map[string]string] // 正确
var s3 Slice[chan int]          // 正确
var s4 Slice[interface{}]       // 正确

type Slice2[T any] []T // 代码等价于 type Slice2[T interface{}] []T
var s5 Slice2[int]

func main() {
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}
