package main

import (
	"fmt"
	"strings"
)

func main() {
	line := "123"
	line = strings.Trim(line, "23")
	fmt.Println(line)

	line2 := "1+2"
	fmt.Printf("line2 is %s\n", line2)
	line2 = strings.Trim(line, "+÷×")
	fmt.Println(line2)

	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))

}
