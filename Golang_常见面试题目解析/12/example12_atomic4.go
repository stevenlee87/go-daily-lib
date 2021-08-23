package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var value int32
	fmt.Println("origin value:", value)

	go entry("1", &value)

	go entry("2", &value)

	time.Sleep(time.Second)
}

func entry(name string, value *int32) {

	swapFlag := atomic.CompareAndSwapInt32(value, 0, 1)

	if swapFlag {
		fmt.Println("goroutine name:", name, ", swap, value:", *value)
	} else {
		fmt.Println("goroutine name:", name, ", not swap, value:", *value)
	}

}
