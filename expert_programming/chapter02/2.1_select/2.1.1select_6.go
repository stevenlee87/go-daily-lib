package main

import "fmt"

// 如果case语句中操作了nil的管道，那么该case语句会被忽略，所以函数会从default出口返回
func main() {
	var c chan string
	select {
	case c <- "hello":
		fmt.Println("sent")
	default:
		fmt.Println("default")
	}
}
