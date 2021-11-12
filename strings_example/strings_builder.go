package main

import (
	"fmt"
	"strings"
)

func main() {
	sArr := []string{"a", "b", "c", "d"}
	var b strings.Builder
	for _, str := range sArr {
		b.WriteString(str) //也可以用fmt.FPrintf(&b,"%s",str)写入其它类型
	}
	fmt.Println(b.String())
}
