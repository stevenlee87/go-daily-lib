package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

/*
Flatten操作即将嵌套很深的字段提到最外层，gabs.Flatten返回一个新的map[string]interface{}，interface{}为 JSON 中叶子节点的值，键为该
叶子的路径。例如：{"foo":[{"bar":"1"},{"bar":"2"}]}执行 flatten 操作之后返回{"foo.0.bar":"1","foo.1.bar":"2"}。
*/
func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{
    "user": {
      "name": "steven",
      "age": 18,
      "members": [
        {
          "name": "hxy",
          "age": 20,
          "relation": "brother1"
        },
        {
          "name": "syf",
          "age": 20,
          "relation": "brother2"
        }
      ],
      "hobbies": ["game", "programming"]
    }
  }`))

	fmt.Println(jObj.StringIndent("", "  "))
	obj, _ := jObj.Flatten()
	fmt.Println(obj)
}
