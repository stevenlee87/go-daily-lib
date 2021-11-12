package main

import "fmt"

func main() {
	s1 := "Hello World!"
	s2 := "你好，中国"

	fmt.Println(len(s1))
	fmt.Println(len(s2))
}

/*
输出
12
15
我们使用len()函数获取到的都是编码后的字节长度，而非字符长度，这一点在使用非 ASCII 字符时很重要
*/
