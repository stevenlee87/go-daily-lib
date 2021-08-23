package main

import (
	"fmt"
)

var t int = 1

var fslice []int

func main() {
	doSomeThing(1, 2, 3, 4)
	return
}

func doSomeThing(arg ...int) { // var arg []int
	fmt.Println(arg) // [1 2 3 4]
	// as a slice
	fslice := []int(arg)
	fmt.Println(fslice[2])
}
