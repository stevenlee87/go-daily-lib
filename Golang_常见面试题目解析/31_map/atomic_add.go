package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	a := int64(0)

	wait := &sync.WaitGroup{}
	wait.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wait.Done()
			atomic.AddInt64(&a, 1)
		}()
	}
	wait.Wait()
	fmt.Println(a)
}
