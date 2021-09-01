package main

import "fmt"

const alphabetStr string = "abcdefghijklmnopqrstuvwxyz"

func main() {
	alphabetMap := make(map[string]bool)
	for _, r := range alphabetStr {
		c := string(r)
		alphabetMap[c] = true
	}
	fmt.Println(alphabetMap["x"])
	alphabetMap["x"] = false
	fmt.Println(alphabetMap["x"])
}
