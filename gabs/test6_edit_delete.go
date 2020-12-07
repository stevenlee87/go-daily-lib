package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	gObj := gabs.New()

	gObj.Set("steven", "info", "name", "first")
	gObj.SetP("lee", "info.name.last")
	gObj.SetJSONPointer(18, "/info/age")

	fmt.Println(gObj.String())

	// 我们也可以调用gabs.Container的StringIndent方法增加前缀和缩进，让输出更美观些：
	fmt.Println(gObj.StringIndent("", "  "))

}
