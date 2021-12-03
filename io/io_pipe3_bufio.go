package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

// 使用 bufio.NewWriter 封装一层 PipeWriter ，使用 bufio.NewScanner 封装一层 PipeReader 实现跟上述代码相同的功能：
func main() {
	r, w := io.Pipe()

	go func() {
		writer := bufio.NewWriter(w)
		for {
			time.Sleep(time.Second)
			writer.WriteString(time.Now().String() + "\n")
			writer.Flush()
		}
	}()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*
writer.Flush() 如果被注释掉的话，会有很长一段时间无数据输出。查看源码发现，其默认的 buffer大小为4096只有当存入的数据大小大于4096 比特时，
它才会调用 PipeWriter.Write() 写入数据。解决方案就是把注释的那行解除注释，每次写入时清空一下缓存就可以了。
*/
