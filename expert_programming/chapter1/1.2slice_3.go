package main

import "fmt"

func SliceExtend() {
	var slice []int
	s1 := append(slice, 1, 2, 3)
	s2 := append(s1, 4)
	fmt.Println(&s1[0] == &s2[0])
	fmt.Println(&s1[0], &s2[0])
}

func main() {
	SliceExtend()
}
