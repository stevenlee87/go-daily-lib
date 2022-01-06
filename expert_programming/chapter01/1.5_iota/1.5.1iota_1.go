package main

import "fmt"

const (
	mutextLocked = 1 << iota // mutext is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

func main() {
	fmt.Println("mutextLocked：", mutextLocked)
	fmt.Println("mutexWoken：", mutexWoken)
	fmt.Println("mutexStarving：", mutexStarving)
	fmt.Println("mutexWaiterShift：", mutexWaiterShift)
	fmt.Println("starvationThresholdNs：", starvationThresholdNs)

}
