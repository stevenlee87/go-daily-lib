package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := strings.Count(str, "")
	fmt.Println(s)

	fmt.Print(str[0:1])
}

/*
// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
*/
