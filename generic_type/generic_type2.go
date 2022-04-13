package main

import (
	"fmt"
	"runtime/debug"
)

func init() {
	debug.SetGCPercent(-1)
}

func main() {
	type Slice[T int | float32 | float64] []T
	var a Slice[int] = []int{1, 2, 3}
	fmt.Println(a)
}
