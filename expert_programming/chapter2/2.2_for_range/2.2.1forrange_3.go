package main

import "fmt"

func RangeNilChannel() {
	var c chan string

	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	RangeNilChannel()
}

/*
for-range表达式作用于值为nil的channel时会永久阻塞

*/
