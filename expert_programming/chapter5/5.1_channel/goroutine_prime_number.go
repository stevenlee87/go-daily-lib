package main

import (
	"fmt"
	"time"
)

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
			fmt.Println("fuck", "i:", i)
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			} else {
				fmt.Printf("%d,%d,%v\n", i, prime, in)
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	//fmt.Println(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch // 新出现的素数
		fmt.Println("prime", prime)
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
		time.Sleep(100 * time.Microsecond)
	}
}
