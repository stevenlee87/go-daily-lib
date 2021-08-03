package main

import "fmt"

const (
	One int = iota + 1
	Two
	Three
	Four
	Five
)

func main() {
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)
	fmt.Println(Five)
}
