package main

import (
	"fmt"
	"time"
)

const time_in_seconds = 10

func executeTheCodeThatMightTakeMoreThanTwentySeconds() {
	time.Sleep(time.Second * 20)
}

func main() {
	// 10 second timer.
	timer := time.NewTimer(time.Second * time_in_seconds)
	defer timer.Stop()

	// 等待定时器完成异步任务
	go func() {
		// 阻塞直到定时器完成。时间到了发送消息给通道 timer.C
		// 该协程中没有其他代码执行
		<-timer.C

		// 如果main() 在10秒之前完成，下面代码不会执行
		fmt.Printf("Congratulations! Your %d second timer finished.", time_in_seconds)
	}()

	// 执行一个耗时方法
	executeTheCodeThatMightTakeMoreThanTwentySeconds()

	// main()函数完成，所以定时器结束.
}
