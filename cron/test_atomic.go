package main

import (
	"log"
	"sync/atomic"
	"time"
)

func main() {
	var d int32
	for i := 0; i < 10; i++ {
		atomic.AddInt32(&d, 1)
		log.Printf("%d: hello world\n", d)
		if atomic.LoadInt32(&d) == 1 {
			time.Sleep(2 * time.Second)
		}
	}

}
