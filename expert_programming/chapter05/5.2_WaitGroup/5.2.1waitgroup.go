package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) //设置计数器，数值即为goroutine的个数
	go func() {
		// Do some work
		time.Sleep(1 * time.Second)

		fmt.Println("Goroutine 1 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	go func() {
		//Do some work
		time.Sleep(2 * time.Second)

		fmt.Println("Goroutine 2 finished!")
		wg.Done() //goroutine执行结束后将计数器减1
	}()

	wg.Wait() //主goroutine阻塞等待计数器变为0
	fmt.Printf("All Goroutine finished!")
}

/*
简单的说，上面程序中wg内部维护了一个计数器：
1. 启动goroutine前将计数器通过Add(2)将计数器设置为待启动的goroutine个数。
2. 启动goroutine后，使用Wait()方法阻塞自己，等待计数器变为0。
3. 每个goroutine执行结束通过Done()方法将计数器减1。
4. 计数器变为0后，阻塞的goroutine被唤醒。
其实WaitGroup也可以实现一组goroutine等待另一组goroutine，这有点像玩杂技，很容出错，
如果不了解其实现原理更是如此。实际上，WaitGroup的实现源码非常简单。
*/
