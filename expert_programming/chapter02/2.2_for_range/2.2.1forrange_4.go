package main

import (
	"fmt"
	"time"
)

func RangeTimer() {
	t := time.NewTimer(time.Second)

	for _ = range t.C {
		fmt.Println("hi")
	}
}

func main() {
	RangeTimer()
}
