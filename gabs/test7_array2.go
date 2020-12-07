package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

/* 我们先通过Array/ArrayP分别在路径user.hobbies和user.bugs下创建数组，然后调用ArrayAppend/ArrayAppendP向这两个数组中添加元素。
现在我们应该可以根据方法有无后缀，后缀是什么来区分它接受什么格式的路径了！
*/
func main() {
	gObj := gabs.New()

	arrObj1, _ := gObj.Array("user", "hobbies")
	fmt.Println(arrObj1.String())

	arrObj2, _ := gObj.ArrayP("user.bugs")
	fmt.Println(arrObj2.String())

	gObj.ArrayAppend("game", "user", "hobbies")
	gObj.ArrayAppend("programming", "user", "hobbies")

	gObj.ArrayAppendP("crash", "user.bugs")
	gObj.ArrayAppendP("panic", "user.bugs")
	fmt.Println(gObj.String())
}
