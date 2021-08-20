package main

import (
	"fmt"
	"runtime"
)

// 代码会触发异常吗？请详细说明
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

/*
输出：
panic: hello或者 1

解析：
结果是随机执⾏。golang 在多个 case 可读的时候会公平的选中⼀个执⾏。

*/
