package main

import "fmt"

func isSubset(s1, s2 []int) bool {
	e := make(map[int]int)

	for i, v := range s1 {
		e[v] = i
	}

	for _, v2 := range s2 {
		if _, ok := e[v2]; !ok {
			return false
		}
	}

	return true
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 5}

	fmt.Println(isSubset(s1, s2))
	fmt.Println(s2[1])
}
