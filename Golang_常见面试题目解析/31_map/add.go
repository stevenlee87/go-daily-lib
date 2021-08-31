package main

import (
	"fmt"
	"sync"
)

func main() {
	var a int = 0

	wait := &sync.WaitGroup{}
	wait.Add(100)
	for i := 0; i < 50; i++ {
		go func() {
			defer wait.Done()
			a++
		}()
	}

	for j := 0; j < 50; j++ {
		go func() {
			defer wait.Done()
			a++
		}()
	}

	wait.Wait()
	fmt.Println(a)
}
