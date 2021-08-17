package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var value int32

	fmt.Println("origin value:", value)

	swapFlag := atomic.CompareAndSwapInt32(&value, 0, 1)

	if swapFlag {
		fmt.Println("swap, value:", value)
	} else {
		fmt.Println("not swap, value:", value)
	}

}
