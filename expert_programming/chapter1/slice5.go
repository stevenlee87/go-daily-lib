package main

import (
	"fmt"
	"reflect"
)

func main() {
	baseStr := "Hello World!"
	fmt.Printf("baseStr: %s\n", baseStr)
	fmt.Printf("baseStr type: %s\n", reflect.TypeOf(baseStr))

	netStr := baseStr[0:5]
	fmt.Printf("netStr: %s\n", netStr)
	fmt.Printf("netStr type: %s\n", reflect.TypeOf(netStr))

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := array[5:7:7]
	fmt.Println("len(slice) = ", len(slice))
	fmt.Println("cap(slice) = ", cap(slice))
}
