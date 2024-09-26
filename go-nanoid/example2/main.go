package main

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func main() {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" // 自定义字母表
	id, err := gonanoid.Generate(alphabet, 10)
	if err != nil {
		fmt.Println("Error generating ID:", err)
		return
	}

	fmt.Println("Generated ID with length 10:", id)
}
