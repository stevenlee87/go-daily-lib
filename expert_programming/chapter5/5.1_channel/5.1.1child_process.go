package main

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	// Do some work...
	time.Sleep(time.Second)

	ch <- 1 // 在管道中写入一个元素表示当前协程已结束
}

func main() {
	channels := make([]chan int, 10) // 创建一个包含10个元素的切片，元素类型为channel

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int) // 在切片中放入一个channel
		go Process(channels[i])      // 启动协程，传一个管道用于通信
	}

	for i, ch := range channels { // 遍历切片，等待子协程结束
		<-ch
		fmt.Println("Routine ", i, " quit！")
	}
}

/*
上面的程序通过创建N个channel来管理N个协程，每个协程都有一个channel用于跟父协程通信，
父协程创建完所欲协程后等待所欲协程结束。

在这个例子中，父协程仅仅是等待子协程结束，其实父协程也可以向管道中写入数据通过子协程结束，
这时子协程需要定期地探测管道中是否有消息出现。
*/
