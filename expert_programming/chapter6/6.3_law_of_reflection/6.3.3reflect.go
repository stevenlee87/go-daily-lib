package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)
	fmt.Println("x :", v.Elem().Interface())
}
