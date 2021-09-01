package main

import (
	"fmt"
	"time"
)

func main() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}
	go func() {
		for a := range abc {
			fmt.Println("a: ", a)
		}
	}()
	close(abc) // 其实这个程序是可以执行的。只不过这么写不太好。应该通过select判断来关闭channel。
	// 可以参考example11,那个程序是无法执行的
	fmt.Println("close")
	time.Sleep(time.Second * 100)
}
