package main

import "fmt"

func func2(i int) {
	i := 0
	fmt.Println(i)
}

func main() {
	var i int
	func2(i)
}

/*
no new variables on left side of :=
*/
