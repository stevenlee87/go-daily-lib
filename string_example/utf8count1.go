package main

import "fmt"

func Utf8Count(s string) int {
	var count int
	for range s {
		count++
	}
	return count
}

func main() {
	fmt.Println(Utf8Count("中国")) // 2
}
