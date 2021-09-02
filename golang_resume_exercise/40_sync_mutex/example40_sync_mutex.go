package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("C chain:", chain)
	chain = chain + " --> C"
}

func main() {
	chain = "main"
	A() // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(chain)
}
