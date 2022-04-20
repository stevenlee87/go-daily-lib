package main

import "fmt"

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

var uintInf Uint // 错误。Uint是一般接口，只能用于类型约束，不得用于变量定义

func main() {
	fmt.Println(uintInf)
}
