package main

import (
	"fmt"
	"strings"
)

/*
func Replace(s, old, new string, n int) string
返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
*/

func main() {
	// non-overlapping: "123" repeat 6 times in s
	s := "123lafaldsjglad123lkfasdf123djfal123lkdjga123lksjfla123l"
	old := "123"
	new := "888"

	fmt.Println("non-overlapping: ")
	// n < 0 ,用 new 替换所有匹配上的 old；n=-1:  888lafaldsjglad888lkfasdf888djfal888lkdjga888lksjfla888l
	fmt.Println("n=-1: ", strings.Replace(s, old, new, -1))

	// 不替换任何匹配的 old；n=0:  123lafaldsjglad123lkfasdf123djfal123lkdjga123lksjfla123l
	fmt.Println("n=0: ", strings.Replace(s, old, new, 0))

	// 用 new 替换第一个匹配的 old；n=1:  888lafaldsjglad123lkfasdf123djfal123lkdjga123lksjfla123l
	fmt.Println("n=1: ", strings.Replace(s, old, new, 1))

	// 用 new 替换前 5 个匹配的 old（实际多于 5 个）；n=5:  888lafaldsjglad888lkfasdf888djfal888lkdjga888lksjfla123l
	fmt.Println("n=5: ", strings.Replace(s, old, new, 5))

	// 用 new 替换前 7 个匹配上的 old（实际没那么多）；n=7:  888lafaldsjglad888lkfasdf888djfal888lkdjga888lksjfla888l
	fmt.Println("n=7: ", strings.Replace(s, old, new, 7))

	// overlapping:
	s = "888888888888888888"
	old = "888"
	new = "666"
	fmt.Println("overlapping: ")

	// n < 0 ,用 new 替换所有匹配上的 old；n=-1:  666666666666666666
	fmt.Println("n=-1: ", strings.Replace(s, old, new, -1))

	// 不替换任何匹配的 old；n=0:  888888888888888888
	fmt.Println("n=0: ", strings.Replace(s, old, new, 0))

	// 用 new 替换第一个匹配的 old；n=1:  666888888888888888
	fmt.Println("n=1: ", strings.Replace(s, old, new, 1))

	// 用 new 替换前 5 个匹配的 old（实际多于 5 个）；n=5:  666666666666666888
	fmt.Println("n=5: ", strings.Replace(s, old, new, 5))

	// 用 new 替换前 7 个匹配上的 old（实际没那么多）；n=7:  666666666666666666
	fmt.Println("n=7: ", strings.Replace(s, old, new, 7))
}
