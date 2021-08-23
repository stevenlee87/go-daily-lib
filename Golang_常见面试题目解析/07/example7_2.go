package main

const alphabetStr string = "abcdefghijklmnopqrstuvwxyz"

func main() {
	var alphabetMap map[string]bool
	for _, r := range alphabetStr {
		c := string(r)
		alphabetMap[c] = true
	}
}

// panic: assignment to entry in nil map
// 因为在声明alphabetMap后并未初始化它，所以它的值是nil, 不指向任何内存地址。
// 需要通过make方法分配确定的内存地址.
// https://github.com/kevinyan815/gocookbook/issues/7
