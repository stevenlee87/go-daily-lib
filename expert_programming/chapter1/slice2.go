package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	b := append(a, 3)
	fmt.Printf("cap(a)=%v,cap(b)=%v\n", cap(a), cap(b))
	c := append(b, 4)
	d := append(b, 5)
	fmt.Printf("cap(a)=%v,cap(b)=%v\n", cap(a), cap(b))
	fmt.Println(a, b, c, d)
	fmt.Printf("%p,%p,%p,%p\n", a, b, c, d)
}

/*
第一次append，超出了a的cap范围，分配一个新的newcap为oldcap*2的array，即4；a不变
第二次append，len(b)是3，cap(b)是4，没有超出b的cap范围，b所依赖的array在len(b)的位置追加4，c共用这个array；b不变
第三次append，由于b没有变化，b所依赖的array在len(b)的位置追加5，会覆盖上一步的4
所以：c[3]和d[3]引用的是同一块内存，都是5

另外，如果 d:= append(c,5) 则结果就是 c[3]=4,d[3]=4,d[4]=5因为这一步会新分配array,并将c的数据拷贝过来
*/
