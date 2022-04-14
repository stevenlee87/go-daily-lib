package main

import (
	"fmt"
)

func main() {
	// 先定义个泛型类型 Slice[T]
	type Slice[T int | string | float32 | float64] []T
	var s Slice[int] = []int{1, 2, 3}
	fmt.Println(s)

	// 在map中套一个泛型类型Slice[T]
	type WowMap[T int | string] map[string]Slice[T]

	a := WowMap[string]{"123": Slice[string]{"1"}}
	b := WowMap[int]{"123": Slice[int]{1}}
	fmt.Println(a, b)
}

/*
这个 WowMap 只有一个类型形参 T，所以只要传入T的实参就行了，比如：WowMap[int]
实例化之后的 WowMap[int] 实际的定义可以看成 map[string]Slice[int] 。这里Slice[int]实际上是实例化了 Slice[T] ，
所以套娃式地把Slicet[T]的定义套进去，最终 WowMap的结构就是： map[string][]int 。一个以string为键，以int切片为值的map
*/
