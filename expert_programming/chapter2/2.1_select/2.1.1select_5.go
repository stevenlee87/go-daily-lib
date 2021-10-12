package main

import "fmt"

func SelectExam5() {
	select {}
}

func main() {
	fmt.Println("start")
	SelectExam5()
	fmt.Println("end")
}

/*
https://github.com/cloudnativebooks/cloud-native-go/issues/44
答案是：C函数陷入阻塞

如果main函数里只调用这一个函数，是会触发死锁检测，从而panic。
原因还是整个函数进入了阻塞状态，而且没有条件能够唤醒。

这个题目说的是函数的描述，没有main（）函数，考查的也不是死锁检测。
*/
