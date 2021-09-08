package main

import (
	"fmt"
)

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}
func main() {
	fmt.Println(a)
	go f()
	c <- 0
	print(a)
}
