package main

import "fmt"

func main() {
	// MyMap类型定义了两个类型形参 KEY 和 VALUE。分别为两个形参指定了不同的类型约束
	// 这个泛型类型的名字叫： MyMap[KEY, VALUE]
	type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

	// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
	var a MyMap[string, float64] = map[string]float64{
		"jack_score": 9.6,
		"bob_score":  8.4,
	}

	for k, v := range a {
		fmt.Println(k, v)
	}
}
