package main

import (
	"fmt"
	"strings"
)

func main() {
	line := "+"
	line = strings.Trim(line, "+÷×")
	fmt.Println(line)
}
