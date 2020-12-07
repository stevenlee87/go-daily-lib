package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

/*
获取数组信息
对于类型为数组的值，gabs提供了几组便捷的查询方法。

获取数组大小：ArrayCount/ArrayCountP，不加后缀的方法接受可变参数作为路径，以P为后缀的方法需要传入.分隔的路径；
获取数组某个索引的元素：ArrayElement/ArrayElementP。
*/
func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{
    "user": {
      "name": "dj",
      "age": 18,
      "members": [
        {
          "name": "hxy",
          "age": 20,
          "relation": "son"
        },
        {
          "name": "syf",
          "age": 20,
          "relation": "brother"
        }
      ],
      "hobbies": ["game", "programming"]
    }
  }`))

	cnt, _ := jObj.ArrayCount("user", "members")
	fmt.Println("member count:", cnt)
	cnt, _ = jObj.ArrayCount("user", "hobbies")
	fmt.Println("hobby count:", cnt)

	ele, _ := jObj.ArrayElement(0, "user", "members")
	fmt.Println("first member:", ele)
	ele, _ = jObj.ArrayElement(1, "user", "hobbies")
	fmt.Println("second hobby:", ele)
}
