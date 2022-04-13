package main

import "fmt"

func main() {
	type Slice[T int | float32 | float64] []T

	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T\n", a) //输出：Type Name: Slice[int]

	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T\n", b) //输出：Type Name: Slice[float32]
}
