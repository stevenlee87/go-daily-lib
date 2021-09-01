package main

import (
	"fmt"
)

func main() {
	fmt.Println([...]string{"1"} == [...]string{"1"}) // 数组只能与相同纬度⻓度以及类型的其他数组⽐较

	// invalid operation: []string{...} == []string{...} (slice can only be compared to nil)
	//fmt.Println([]string{"1"} == []string{"1"}) // 切⽚之间不能直接⽐较。。

}
