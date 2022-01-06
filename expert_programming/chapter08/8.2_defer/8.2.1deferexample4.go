package main

import "fmt"

func DeferDemo5() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}

func main() {
	fmt.Println(DeferDemo5())
}

/*
OutPut:
2

函数的return语句并不是原子的，实际执行分为设置返回值和ret两步，defer语句实际
执行在返回前，即拥有defer的函数返回过程是：设置返回值->执行defer->ret。所以return
语句先把result设置为i的值，即1，defer语句中又把result递增1，所以最终返回2。
*/
