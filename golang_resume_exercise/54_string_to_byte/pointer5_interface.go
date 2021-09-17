package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := (*interface{})(nil)
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a))
	var b interface{} = (*interface{})(nil)
	fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b))
	fmt.Println(a == nil, b == nil)
}
