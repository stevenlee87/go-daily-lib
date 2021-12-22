package main

import "fmt"

func foo1() {
	foo := func() int {
		defer func() {
			recover()
		}()

		panic("demo")
		return 10
	}
	ret := foo()
	fmt.Println(ret)
}

func main() {
	foo1()
}

/*
Output:
0
*/
