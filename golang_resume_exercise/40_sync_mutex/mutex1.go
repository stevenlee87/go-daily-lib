package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

func main() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"

	mu.Lock()
	defer mu.Unlock()
	fmt.Println("C chain:", chain)
	chain = chain + " --> C"
}
