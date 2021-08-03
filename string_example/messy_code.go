package main

import "fmt"

func main() {
	s := "中国"
	fmt.Println(s[:5])

	b := []byte{129, 130, 131}
	fmt.Println(string(b))
}

/*
上面输出：

中��
���
因为“国”编码有 3 个字节，s[:5]只取了前两个，这两个字节无法组成一个合法的 UTF8 字符，故输出两个�。
*/
