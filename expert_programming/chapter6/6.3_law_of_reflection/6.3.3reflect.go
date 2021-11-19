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

// 该例演示了反射的第三个能力，即通过反射可以修改interface变量，但必须获得interface变量的地址
