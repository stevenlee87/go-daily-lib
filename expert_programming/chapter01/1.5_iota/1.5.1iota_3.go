package main

import "fmt"

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // iota :1<<0=1  1 - 1 = 0
	bit1, mask1
	_, _
	bit3, mask3
)

func main() {
	fmt.Printf("bit0:%v, mask0:%v\n", bit0, mask0)
	fmt.Printf("bit1:%v, mask1:%v\n", bit1, mask1)
	fmt.Printf("bit3:%v, mask3:%v\n", bit3, mask3)

}
