package main

import "fmt"

/*
在 Go 中，字符串中的每个字符都是以 Unicode 编码的整数表示的。
当你使用 s[0] 访问字符串 s 中的第一个字符时，它实际上返回的是字符 'a' 的 Unicode 编码，而不是字符本身。
字符 'a' 的 Unicode 编码是 97。因此，你看到的打印结果是 97，而不是字符 'a'。

如果你希望打印字符 'a' 而不是其 Unicode 编码，可以使用 %c 格式化占位符，如下所示：
*/
func main() {
	m := map[byte]int{}
	var s string
	s = "abca"
	fmt.Println(s[0])
	fmt.Printf("s[0]:%c\n", s[0]) // 打印字符 'a'
	fmt.Println(m[s[0]])
	m[s[0]]++

	fmt.Printf("s[3]: %c\n", s[3]) // 打印字符s[3]
	fmt.Println(m[s[3]])

	delete(m, s[0])
	fmt.Println(m[s[3]])
}
