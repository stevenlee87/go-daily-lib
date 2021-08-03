package main

import "fmt"

const (
	Mask1 int = 1<<(iota+1) - 1 // 2^1 - 1 = 1
	Mask2                       // 2^2 - 1 = 3
	Mask3
	Mask4
)

func main() {
	fmt.Println(Mask1)
	fmt.Println(Mask2)
	fmt.Println(Mask3)
	fmt.Println(Mask4)

}

/*
1 << 1 = 2
1 << 2 = 4
1 << 3 = 8
1 << 4 = 16
1 << 5 = 32
1 << 6 = 64
1 << 7 = 128
1 << 8 = 256
1 << 9 = 512

512 >> 1 = 256
512 >> 2 = 128
512 >> 3 = 64
512 >> 4 = 32
512 >> 5 = 16
512 >> 6 = 8
512 >> 7 = 4
512 >> 8 = 2
512 >> 9 = 1
*/
