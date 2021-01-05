package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

/*
函数
如果仅仅能进行常规的算数和逻辑运算，govaluate的功能会大打折扣。govaluate提供了自定义函数的功能。
所有自定义函数需要先定义好，存入一个map[string]govaluate.ExpressionFunction变量中，
然后调用govaluate.NewEvaluableExpressionWithFunctions()生成表达式，此表达式中就可以使用这些函数了。
自定义函数类型为func (args ...interface{}) (interface{}, error)，如果函数返回错误，则这个表达式求值返回错误。
*/

func main() {
	functions := map[string]govaluate.ExpressionFunction{
		"strlen": func(args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return length, nil
		},
	}

	exprString := "strlen('teststring')"
	expr, _ := govaluate.NewEvaluableExpressionWithFunctions(exprString, functions)
	result, _ := expr.Evaluate(nil)
	fmt.Println(result)
}

/*
上面例子中，我们定义一个函数strlen计算第一个参数的字符串长度。表达式strlen('teststring')调用strlen函数返回字符串teststring的长度。
函数可以接受任意数量的参数，而且可以处理嵌套函数调用的问题。所以可以写出类似下面这种复杂的表达式：

sqrt(x1 ** y1, x2 ** y2)
max(someValue, abs(anotherValue), 10 * lastValue)
*/
