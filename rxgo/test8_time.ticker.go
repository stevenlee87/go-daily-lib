package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(5 * time.Second)

	var count int
	for range t.C {
		fmt.Println(count)
		count++
	}
}
