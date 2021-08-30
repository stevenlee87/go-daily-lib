package main

import "fmt"

func Add(arr []int, index int, value int) []int {
	if len(arr) == index {
		return append(arr, value)
	}
	fmt.Println("arr[:3] is:", arr[:3])
	arr = append(arr[:index+1], arr[index:]...) // arr = append(arr[:3], arr[2:])
	arr[index] = value
	return arr
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(Add(arr, 2, 5))
}
