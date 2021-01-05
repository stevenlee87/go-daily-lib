package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

/*
在 Go 语言中，访问器（Accessors）就是通过.操作访问结构中的字段。
如果传入的参数中有结构体类型，govaluate也支持使用.访问其内部字段或调用它们的方法：
*/

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func (u User) Fullname() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	u := User{FirstName: "li", LastName: "dajun", Age: 18}
	fmt.Println(u)
	parameters := make(map[string]interface{})
	parameters["u"] = u

	//expr, _ := govaluate.NewEvaluableExpression("u.Fullname()")
	expr, _ := govaluate.NewEvaluableExpression("u")
	result, _ := expr.Evaluate(parameters)
	fmt.Println("user", result)

	//expr, _ = govaluate.NewEvaluableExpression("u.Age > 18")
	//result, _ = expr.Evaluate(parameters)
	//fmt.Println("age > 18?", result)
}

/*
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10d350f]

*/
