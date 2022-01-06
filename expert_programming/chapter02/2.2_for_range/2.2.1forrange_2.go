package main

import (
	"fmt"
	"sync"
)

func PrintSlice() {
	s := []int{1, 2, 3}
	var wg sync.WaitGroup

	wg.Add(len(s))
	for _, v := range s {
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	PrintSlice()
}

/*
在for-range表达式中，迭代变量只会声明一次，在多次迭代中共享。上面的程序中由于循环迭代速度
快于协程启动速度，协程最终访问的v值为3，所以最终输出3个3
*/
