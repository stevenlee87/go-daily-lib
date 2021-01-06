package main

import (
	"fmt"
	"log"

	"github.com/Knetic/govaluate"
)

func main() {
	exprString := `>>>`
	expr, err := govaluate.NewEvaluableExpression(exprString)
	if err != nil {
		log.Fatal("syntax error:", err)
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		log.Fatal("evaluate error:", err)
	}
	fmt.Println(result)
}

// 2021/01/06 10:58:43 syntax error:Invalid token: '>>>'
