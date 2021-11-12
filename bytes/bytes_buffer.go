package main

import (
	"bytes"
	"fmt"
)

func main() {
	sArr := []string{"a", "b", "c", "d"}
	var buffer bytes.Buffer
	for _, str := range sArr {
		//fmt.Println(str)
		buffer.WriteString(str)
	}
	fmt.Println(buffer.String())
}
