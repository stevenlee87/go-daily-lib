package main

import (
	"fmt"
	"unsafe"
)

func main() {
	mp := make(map[string]int)
	mp["steven"] = 100
	mp["shifu"] = 18
	//mp["jack"] = 60

	// type ArbitraryType int
	// type Pointer *ArbitraryType
	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}

/*
count 的转换过程：
&mp => pointer => **int => int
*/
