package main

import "fmt"

func main() {
	var a int = 2
	fmt.Println(&a)
	a = 3
	fmt.Println(&a)
	fmt.Println(a)
}
