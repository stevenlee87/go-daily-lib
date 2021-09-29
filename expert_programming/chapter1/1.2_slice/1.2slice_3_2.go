package main

import "fmt"

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func main() {
	var slice []int
	slice = append(slice, 1, 2, 3)

	newSlice := AddElement(slice, 4)
	fmt.Println(&slice[0] == &newSlice[0])
}
