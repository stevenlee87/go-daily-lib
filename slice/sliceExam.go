package main

import "fmt"

func main() {
	numbers := make([]int, 2, 5)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
