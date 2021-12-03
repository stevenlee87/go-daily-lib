package main

import (
	"fmt"
	"io"
	"time"
)

// 这里写一个简单的例子，在 goroutine 中，每隔 1 秒向 PipeWriter 中写入当前的时间，从 PipeReader 读取写入的数据。不使用任何工具函数
func main() {
	r, w := io.Pipe()

	go func() {
		for {
			time.Sleep(time.Second)
			w.Write([]byte(time.Now().String()))
		}
	}()
	for {
		dataRead := make([]byte, 256)
		n, _ := r.Read(dataRead)
		fmt.Println(string(dataRead[:n]))
	}
}

/*
如果阅读其源代码，它内部有一个 pipe 对象，将其 Read 相关功能和 Write 相关功能拆分成了 PipeReader 和 PipeWriter 暴露出去。

PipeWriter.Write()方法写入数据，PipeReader.Read()读取数据。由于其实现了接口 io.ReadCloser 和 io.WriteCloser 。
所以可以使用更高层的工具函数来操作它们。

对于 PipeReader ，可以使用 bytes.Buffer.ReadFrom 或者 bufio.NewScanner 。

对于 PipeWriter 可以使用 fmt.Fprintf() 或者 bufio.NewWriter
*/
