package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := array[3:][:3:3]
	fmt.Println(s)
}
