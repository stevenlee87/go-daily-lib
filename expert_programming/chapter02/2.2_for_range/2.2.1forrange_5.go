package main

import "fmt"

func RangeDemo() {
	s := []int{1, 2, 3}

	for i := range s {
		s = append(s, i)
	}
	fmt.Println(s)
}

func main() {
	RangeDemo()
}
