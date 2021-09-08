package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	count int
	sync.Mutex
}

func main() {
	var mu MyMutex
	mu.Lock()
	var mu2 = mu
	mu.count++
	mu.Unlock()

	//mu2.Unlock()
	//var mu2 = mu
	mu2.Mutex = sync.Mutex{} // Mutex是一个互斥的排他锁，零值Mutex为未上锁状态
	mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)

}

/* fatal error: all goroutines are asleep - deadlock!
加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁。
*/
