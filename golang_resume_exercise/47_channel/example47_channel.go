package main

import "fmt"

func main() {
	var ch chan int
	var count int

	go func() {
		ch <- 1
	}()

	go func() {
		count++
		close(ch) // ch 未有被初始化，关闭时会报错。
	}()

	<-ch
	fmt.Println(count)
}

// panic: close of nil channel
