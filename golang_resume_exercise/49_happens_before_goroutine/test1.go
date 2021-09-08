package main

import "fmt"

func add(a int) {
	a = 3
	fmt.Println(a)
	fmt.Println("edit", &a)
}

func main() {
	var a int = 2
	fmt.Println("init", &a)
	add(a)
	fmt.Println(a)
	fmt.Println("init", &a)
}
