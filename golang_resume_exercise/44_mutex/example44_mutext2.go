package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	lock  *sync.Mutex
}

func main() {
	mu := MyMutex{
		count: 0,
		lock:  &sync.Mutex{},
	}

	mu.lock.Lock()
	var mu2 = mu
	mu.count++
	mu.lock.Unlock()

	mu2.lock.Lock()
	mu2.count++
	mu2.lock.Unlock()
	fmt.Println(mu.count, mu2.count)

}
