package main

import "fmt"

// 我们也可以以字符形式输出：
func main() {
	s := "Go 语言"

	for index, c := range s {
		fmt.Printf("%d %c\n", index, c)
	}
}

/*
0 G
1 o
2
3 语
6 言
*/
