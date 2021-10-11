package main

import (
	"fmt"
)

func main() {
	s := "中国"
	for index, value := range s {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}

	// 字符串的长度是指字节数，而非字符数，比如汉字"中"和"国"的UTF-8编码各占
	// 3个字节，字符串"中国"的长度为6而不是2
	fmt.Println(len(s))
	fmt.Printf("\u0000") // \u0000是空字符
	fmt.Printf("123\n")
	fmt.Println(len("\u0000"))
	fmt.Println(len(""))
}

/*
Unicode编码中\u0000代表的是空字符，属于控制字符，也叫不可显字符。

这个空字符与空格不同，空格的编号是 \u0020

"\u0000"就是一个包含一个字符的字符串

""是一个空字符串，这个字符串中一个字符也没有，所以是0
*/
