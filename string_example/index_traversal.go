package main

import "fmt"

func main() {
	s := "Go 语言"

	for index, c := range s {
		fmt.Println(index, c)
	}
}

/*
0 71
1 111
2 32
3 35821
6 35328
*/
