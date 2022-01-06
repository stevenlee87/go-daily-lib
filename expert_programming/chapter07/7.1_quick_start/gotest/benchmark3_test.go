package gotest_test

import (
	"testing"

	"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest"
)

func BenchmarkMakeByErrorsNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeByErrorsNew()
	}
}

func BenchmarkMakeByFmtErrorf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeByFmtErrorf()
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkMakeByErrorsNew
BenchmarkMakeByErrorsNew-8   	1000000000	         0.3056 ns/op
BenchmarkMakeByFmtErrorf
BenchmarkMakeByFmtErrorf-8   	13585872	        83.59 ns/op
PASS
*/
