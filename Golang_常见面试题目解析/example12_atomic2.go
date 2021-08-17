package main

import (
	"fmt"
	"sync/atomic"
)

var value int32

func SetValue(delta int32) {
	v := value
	if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
		fmt.Println("ok")
	}
}

func main() {
	SetValue(1)
}

// atomic.CompareAndSwapInt32 函数不需要循环调⽤。
