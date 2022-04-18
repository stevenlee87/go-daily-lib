package main

import "fmt"

func Add[T int | float32 | float64](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add[int](2, 3))

	fmt.Println(Add(4, 5))

	fmt.Println(Add(1.0, 2.0))
}
