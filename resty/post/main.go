package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func main() {
	// 创建一个 Resty 对象
	client := resty.New()

	// 发送 POST 请求
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"name":"john", "age":30}`).
		Post("https://httpbin.org/post")

	// 处理响应
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.String())

	fmt.Println(resp.StatusCode())
}
