package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

/*
支持的操作和类型
govaluate支持的操作和类型与 Go 语言有些不同。一方面govaluate中的类型和操作不如 Go 丰富，另一方面govaluate也对一些操作进行了扩展。

算数、比较和逻辑运算：
+-/*&|^**%>><<：加减乘除，按位与，按位或，异或，乘方，取模，左移和右移；
>>=<<===!==~!~：=~为正则匹配，!~为正则不匹配；
||&&：逻辑或和逻辑与。


常量：
数字常量，govaluate中将数字都作为 64 位浮点数处理；
字符串常量，注意在govaluate中，字符串用单引号'；
日期时间常量，格式与字符串相同，govaluate会尝试自动解析字符串是否是日期，只支持 RFC3339、ISO8601等有限的格式；
布尔常量：true、false。
其他：

圆括号可以改变计算优先级；
数组定义在()中，每个元素之间用,分隔，可以支持任意的元素类型，如(1, 2, 'foo')。实际上在govaluate中数组是用[]interface{}来表示的；
三目运算符：? :。
在下面代码中，govaluate会先将2020-01-06和2020-01-05 23:59:59转为time.Time类型，然后再比较大小：
*/

func main() {
	expr, _ := govaluate.NewEvaluableExpression("'2020-01-06' > '2020-01-05 23:59:59'")
	result, _ := expr.Evaluate(nil)
	fmt.Println(result)
}
