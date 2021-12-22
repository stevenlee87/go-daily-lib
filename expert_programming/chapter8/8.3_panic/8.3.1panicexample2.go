package main

import "fmt"

func foo() {
	defer fmt.Print("A")
	defer fmt.Print("B")

	fmt.Print("C")
	panic("demo")
	defer fmt.Print("D")
}

func main() {
	defer func() {
		recover()
	}()

	defer func() {
		fmt.Print("1")
	}()

	foo()
}

/*
Output:
CBA1

panic会触发本函数中所有的defer函数，然后将异常抛给上层调用函数，继续触发上层调用函数的defer函数，main()最终输出CBA1
*/
