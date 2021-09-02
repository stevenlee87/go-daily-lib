package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	letter, number := make(chan bool), make(chan bool)
	var wait sync.WaitGroup

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()

	wait.Add(1)
	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}

				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}()
	number <- true
	wait.Wait()

	fmt.Println()
	costTime := time.Since(start)
	fmt.Println(costTime)
}
