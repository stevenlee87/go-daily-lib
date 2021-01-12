package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//拼凑json   body为map数组
	var rbody []map[string]interface{}
	t1 := make(map[string]interface{})
	t1["id"] = 1
	t1["name"] = "Python"
	t1["catagory"] = 1

	rbody = append(rbody, t1)
	t2 := make(map[string]interface{})
	t2["id"] = 1
	t2["name"] = "Linux基础"
	t2["catagory"] = 4

	rbody = append(rbody, t2)

	cnnJson := make(map[string]interface{})
	cnnJson["error_no"] = 0
	cnnJson["data"] = rbody

	b, _ := json.Marshal(cnnJson)

	cnnn := string(b)
	fmt.Println(cnnn)

	r := gin.Default()
	r.GET("/category", func(c *gin.Context) {
		c.JSON(200, cnnJson)
	})
	r.Run(":8080")
}
