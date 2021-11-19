package main

import (
	"fmt"
	"reflect"
)

func main() {
	var A interface{}
	A = 100

	v := reflect.ValueOf(A)

	B := v.Interface()

	if A == B { // true
		fmt.Println("They are same!\n")
	}
}

/* 在上面的函数中，通过relect.Value()获取接口变量A的反射对象，然后又通过反射对象的Interface()获取B，
结果A和B是相同的。

这个例子展示了反射的第二个能力，即能够从一个反射对象还原到原来的interface对象。
*/
