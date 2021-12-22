package main

import "fmt"

func main() {
	defer func() {
		func() { // recover 在defer嵌套函数中无效
			if err := recover(); err != nil {
				fmt.Println("A")
			}
		}()
	}()

	panic("demo")
	fmt.Println("B")
}

/*
Output:
panic: demo

goroutine 1 [running]:
main.main()
*/
