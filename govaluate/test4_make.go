package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

/*
一次“编译”多次运行
使用带参数的表达式，我们可以实现一个表达式的一次“编译”，多次运行。只需要使用编译返回的表达式对象即可，可多次调用其Evaluate()方法：
*/

func main() {
	expr, _ := govaluate.NewEvaluableExpression("a + b")
	parameters := make(map[string]interface{})
	parameters["a"] = 1
	parameters["b"] = 2
	result, _ := expr.Evaluate(parameters)
	fmt.Println(result)

	parameters = make(map[string]interface{})
	parameters["a"] = 10
	parameters["b"] = 20
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)
}
