package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

/*
命名
使用govaluate与直接编写 Go 代码不同，在 Go 代码中标识符中不能出现-、+、$等符号。govaluate可以通过转义使用这些符号。有两种转义方式：

将名称用[和]包裹起来，例如[response-time]；
使用\将紧接着下一个的字符转义。
例如：
*/
func main() {
	expr, _ := govaluate.NewEvaluableExpression("[response-time] < 100")
	parameters := make(map[string]interface{})
	parameters["response-time"] = 80
	result, _ := expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression("response\\-time < 100")
	parameters = make(map[string]interface{})
	parameters["response-time"] = 80
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)
}

/*
注意一点，因为在字符串中\本身就是需要转义的，所以在第二个表达式中要使用\\。或者可以使用
`response\-time` < 100
*/
