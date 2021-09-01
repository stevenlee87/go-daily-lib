package main

import (
	"fmt"
	"sync"
	"time"
)

/*
为 sync.WaitGroup 中Wait函数⽀持 WaitTimeout 功能.

⾸先 sync.WaitGroup 对象的 Wait 函数本身是阻塞的，同时，超时⽤到的 time.Timer 对象也需要阻塞的读。
同时阻塞的两个对象肯定要每个启动⼀个协程,每个协程去处理⼀个阻塞，难点在于怎么知道哪个阻塞先完成。
⽬前我⽤的⽅式是声明⼀个没有缓冲的 chan ，谁先完成谁优先向管道中写⼊数据。
*/

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	ch := make(chan bool, 1)
	go time.AfterFunc(timeout, func() {
		ch <- true
	})

	go func() {
		wg.Wait() // 阻塞的！！！
		ch <- false
	}()

	return <-ch
}

func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close // 阻塞的！！！
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}
