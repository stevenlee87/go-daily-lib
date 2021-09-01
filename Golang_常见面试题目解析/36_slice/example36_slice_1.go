package main

import (
	"fmt"
)

func main() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:] // [b c]
	str2[1] = "new"  // [b new]
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
	fmt.Println(str2)
}

/*
[a, b, new]
[a, b, new]
[b, new, z, x, y]

golang 中的切⽚底层其实使⽤的是数组。当使⽤ str1[1:] 使， str2 和 str1 底层共享⼀个数组，
这会导致 str2[1] = "new" 语句影响 str1 。

⽽ append 会导致底层数组扩容，⽣成新的数组，因此追加数据后的 str2 不会影响str1 。
但是为什么对 str2 复制后影响的却是 str1 的第三个元素呢？这是因为切⽚ str2 是从数组的第⼆个元素开始，
str2 索引为 1 的元素对应的是 str1 索引为2 的元素。
*/
