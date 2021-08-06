package main

import (
	"fmt"
	"sync"
	"time"
)

/*
waitgroup在go中，用于线程同步，指等待一个组，等待一个系列执行完成后，才会向下执行
*/
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calc(&wg, i)
	}
	wg.Wait()
	fmt.Println("all goroutine finish")
}
func calc(w *sync.WaitGroup, i int) {
	fmt.Println("calc:", i)
	time.Sleep(time.Second)
	w.Done()
}
