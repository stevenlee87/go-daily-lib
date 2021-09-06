package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		//wg.Add(1) 注释掉这行程序就正常了。
	}()
	wg.Wait()
}

// 当你Add()之前，就Wait()了，就会发生这个错误。
