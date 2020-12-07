package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

/*
存在性判断
gabs提供了两个方法检查对应的路径上是否存在数据：

Exists(hierarchy ...string)；
ExistsP(path string)：方法名以P结尾，表示接受以.分隔的路径。
*/
func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{"user":{"name": "steven","age": 18}}`))
	fmt.Printf("has name? %t\n", jObj.Exists("user", "name"))
	fmt.Printf("has age? %t\n", jObj.ExistsP("user.age"))
	fmt.Printf("has job? %t\n", jObj.Exists("user", "job"))
}
