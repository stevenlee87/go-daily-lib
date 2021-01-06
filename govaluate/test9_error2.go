package main

import (
	"fmt"
	"log"

	"github.com/Knetic/govaluate"
)

func main() {
	exprString := `foo > 0`
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

/*
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x10d29ce]
*/
