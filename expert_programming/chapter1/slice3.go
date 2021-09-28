package main

import "fmt"

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		fmt.Println("i is:", i)
		s[i]++
	}
	fmt.Printf("%p, %v\n", s, s)
}

func main() {
	s1 := []int{1, 2}
	s2 := make([]int, 2, 3)
	copy(s2, s1)
	s2 = append(s2, 3)
	fmt.Println("cap(s2) is:", cap(s2))

	fmt.Println(s1, s2)
	SliceRise(s1)
	fmt.Println(s1, s2)
	fmt.Printf("s2: %p, %v\n", s2, s2)
	SliceRise(s2)
	fmt.Println(s1, s2)
}
