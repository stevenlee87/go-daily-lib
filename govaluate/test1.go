package main

import (
	"fmt"
	"log"

	"github.com/Knetic/govaluate"
)

func main() {
	expr, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		log.Fatal("syntax error:", err)
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		log.Fatal("evaluate error:", err)
	}

	fmt.Println(result)
}

/*
使用govaluate计算表达式只需要两步：
1.调用NewEvaluableExpression()将表达式转为一个表达式对象；
2.调用表达式对象的Evaluate方法，传入参数，返回表达式的值。

上面演示了一个很简单的例子，我们使用govaluate计算10 > 0的值，该表达式不需要参数，故传给Evaluate()方法nil值。
当然，这个例子并不实用，显然我们直接在代码中计算10 > 0更简单。
但问题是，有些时候我们并不知道需要计算的表达式的所有信息，甚至我们都不知道表达式的结构。这时govaluate的作用就体现出来了。
*/
