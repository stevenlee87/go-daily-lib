package main

import "fmt"

func main() {
	var big int64 = 129
	var little = int8(big)
	fmt.Println(little)

	var a map[string]string
	a = make(map[string]string, 10)
	fmt.Println(a)

}

/*
先把129转换成2进制1000 0001,然后十进制就是-1
再把-1转换成补码：
原码：0000 0001
反码：1111 1110
补码：1111 1111
也就是-127
*/
