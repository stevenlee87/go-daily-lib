package main

import "fmt"

func main() {
	sl := make([]int, 0, 10)
	var appenFunc = func(s []int) {
		s = append(s, 10, 20, 30)
		fmt.Println(s)
	}
	fmt.Println(sl)
	appenFunc(sl)
	fmt.Println(sl)
	fmt.Println(sl[:10])
}

/*
[]
[10 20 30]
[]
[10 20 30 0 0 0 0 0 0 0]
*/
