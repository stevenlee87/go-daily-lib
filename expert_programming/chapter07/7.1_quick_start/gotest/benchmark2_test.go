package gotest_test

import (
	"testing"
	"time"
)

func BenchmarkSetBytes(b *testing.B) {
	b.SetBytes(1024 * 1024)
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second) // 模拟待测函数
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkSetBytes
BenchmarkSetBytes-8   	       1	1004217211 ns/op	   1.04 MB/s
PASS

*/
