package main

import "fmt"

func main() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("len(s1) = %d\n", len(s1))
	fmt.Printf("len(s2) = %d\n", len(s2))

	fmt.Printf("cap(s1) = %d\n", cap(s1))
	fmt.Printf("cap(s2) = %d\n", cap(s2))

	fmt.Printf("---------------------------------------\n")

	s1 = append(s1, 0)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("len(s1) = %d\n", len(s1))
	fmt.Printf("len(s2) = %d\n", len(s2))

	fmt.Printf("cap(s1) = %d\n", cap(s1))
	fmt.Printf("cap(s2) = %d\n", cap(s2))
}
