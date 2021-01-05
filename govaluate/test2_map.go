package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, _ := govaluate.NewEvaluableExpression("foo > 0")
	parameters := make(map[string]interface{})
	parameters["foo"] = -1
	result, _ := expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")
	parameters = make(map[string]interface{})
	parameters["requests_made"] = 100
	parameters["requests_succeeded"] = 80
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)

	expr, _ = govaluate.NewEvaluableExpression("(mem_used / total_mem) * 100")
	parameters = make(map[string]interface{})
	parameters["total_mem"] = 1024
	parameters["mem_used"] = 512
	result, _ = expr.Evaluate(parameters)
	fmt.Println(result)
}

/*
第一个表达式中，我们想要计算foo > 0的结果，在传入参数中将foo设置为 -1，最终输出false。

第二个表达式中，我们想要计算(requests_made * requests_succeeded / 100) >= 90的值，
在参数中设置requests_made为 100，requests_succeeded为 80，结果为true。

上面两个表达式都返回bool结果，第三个表达式返回一个浮点数。(mem_used / total_mem) * 100
根据传入的总内存total_mem和当前使用内存mem_used，返回内存占用百分比，结果为 50。
*/
