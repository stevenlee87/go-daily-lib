# Chapter 7 测试

目前go test支持的测试类型有：
- 单元测试
- 性能测试
- 示例测试

## 7.1 快速开始

### 1.单元测试  
单元测试是指对软件中的最小测试单元进行检测和验证，比如对一个函数的测试。

### 2.性能测试  
性能测试也成为基准测试，可以测试一段程序的性能，可以得到时间消耗、内存使用情况的报告。

### 3.示例测试  
广泛应用于Go源码和各种开源框架中，用于展示某个包或某个方法的用法。  
比如，Go标准库中，mail包展示如何从一个字符串解析出邮件列表的用法，非常直观易懂。  
源码位于 src/net/mail/example_test.go 中：  
```go
func ExampleParseAddressList() {
    const list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
    emails, err := mail.ParseAddressList(list)
    if err != nil {
        log.Fatal(err)
    }

    for _, v := range emails {
        fmt.Println(v.Name, v.Address)
    }

    // Output:
    // Alice alice@example.com
    // Bob bob@example.com
    // Eve eve@example.com
}
```

## 总结
### 1. 单元测试  
通过package语句可以看到，测试文件属于“gotest_test”包，测试文件也可以跟源文件在同一个包，但常见的做法是创建一个包专用于测试，
这样可以使测试文件和源文件隔离。GO源代码以及其他知名的开源框架通常会创建测试包，而且规则是在原包名上加上”_test”。

测试函数命名规则为”TestXxx”，其中“Test”为单元测试的固定开头，go test只会执行以此为开头的方法。
紧跟“Test”是以首字母大写的单词，用于识别待测试函数。

测试函数参数并不是必须要使用的，但”testing.T”提供了丰富的方法帮助控制测试流程。

t.Errorf()用于标记测试失败，标记失败还有几个方法，在介绍testing.T结构时再详细介绍。 

go test

### 2. 性能测试
性能测试函数命名规则为”BenchmarkXxx”，其中”Xxx”为自定义的标识，需要以大写字母开始，通常为待测函数。

testing.B提供了一系列的用于辅助性能测试的方法或成员，比如本例中的 b.N 表示循环执行的次数，而N值不用程序员特别关心，按照官方说法，
N值是动态调整的，直到可靠的算出程序执行时间后才会停止，具体执行次数会在执行结束后打印出来。

其中 -bench 为go test的flag，该flag指示go test进行性能测试，即执行测试文件中符合”BenchmarkXxx”规 则的方法。
紧跟flag的为flag的参数，本例表示执行当前所有的性能测试。

go test -bench=.

### 3. 示例测试
1. 例子测试函数名需要以”Example”开头；
2. 检测单行输出格式为“// Output: <期望字符串>”；
3. 检测多行输出格式为“// Output: \ <期望字符串> \ <期望字符串>”，每个期望字符串占一行；
4. 检测无序输出格式为”// Unordered output: \ <期望字符串> \ <期望字符串>”，每个期望字符串占一 行；
5. 测试字符串时会自动忽略字符串前后的空白字符；
6. 如果测试函数中没有“Output”标识，则该测试函数不会被执行；
7. 执行测试可以使用 go test ，此时该目录下的其他测试文件也会一并执行；
8. 执行测试可以使用 go test <xxx_test.go> ，此时仅执行特定文件中的测试函数；

go test example_test.go

## 7.2 进阶测试

### 7.2.1 子测试

简单的说，子测试提供一种在一个测试函数中执行多个测试的能力，比如原来有TestA、TestB和TestC三个测试函数，
每个测试函数执行开始都需要做些相同的初始化工作，那么可以利用子测试将这三个测试合并到一个测试中，
这样初始化工作只需要做一次。 除此之外，子测试还提供了诸多便利，下面我们逐一说明。

go test subunit_test.go -v

#### 子测试命名规则  
通过上面的例子我们知道 Run() 方法第一个参数为子测试的名字，而实际上子测试的内部命名规则为：”<父测试名 字>/<传递给Run的名字>“。
比如，传递给 Run() 的名字是“A=1”，那么子测试名字为“TestSub/A=1”。这个在上面的命令行输出中也可以看出。

#### 过滤筛选  
通过测试的名字，可以在执行中过滤掉一部分测试。 比如，只执行上例中“A=*”的子测试，那么执行时使用 -run Sub/A= 参数即可；

#### 子测试并发  
前面提到的多个子测试共享setup和teardown有一个前提是子测试没有并发，如果子测试使用 t.Parallel() 指定并发，那么就没办法共享teardown了，
因为执行顺序很可能是setup->子测试1->teardown->子测试2…。 
如果子测试可能并发，则可以把子测试通过 Run() 再嵌套一层，Run() 可以保证其下的所有子测试执行结束后再返回。

go test subparallel_test.go -v -run SubParallel

## 总结
- 子测试适用于单元测试和性能测试； 
- 子测试可以控制并发； 
- 子测试提供一种类似table-driven风格的测试； 
- 子测试可以共享setup和tear-down；  
注释：setup 和 teardown 是在每个 case 执行前后都需要执行的操作

### 7.2.1 Main测试

我们知道子测试的一个方便之处在于可以让多个测试共享Setup和Tear-down。但这种程度的共享有时并不满足需求，
有时希望在整个测试程序做一些全局的setup和Tear-down，这时就需要Main测试了。 

所谓Main测试，即声明一个 func TestMain(m *testing.M) ，它是名字比较特殊的测试，参数类型为 testing.M 指针。
如果声明了这样一个函数，当前测试程序将不是直接执行各项测试，而是将测试交给TestMain调度。

## 7.3 实现原理

### 7.3.1 testing.common

我们知道单元测试函数需要传递一个 testing.T 类型的参数，而性能测试函数需要传递一个 testing.B 类型的参 数，该参数可用于控制测试的流程，
比如标记测试失败等。 

testing.T 和 testing.B 属于 testing 包中的两个数据类型，该类型提供一系列的方法用于控制函数执行流程， 考虑到二者有一定的相似性，
所以Go实现时抽象出一个 testing.common 作为一个基础类型， 而 testing.T 和 testing.B 则属于 testing.common 的扩展。

common.helpers 已经被取消了：https://github.com/golang/go/commit/4c174a7ba66724f8f9a1915c8f4868a8b3aaf219

### 7.3.2 testing.TB接口

TB接口，顾名思义，是testing.T(单元测试)和testing.B(性能测试)共用的接口。 

TB接口通过在接口中定义一个名为private(）的私有方法，保证了即使用户实现了类似的接口，也不会跟 testing.TB接口冲突。 

其实，这些接口在testing.T和testing.B公共成员testing.common中已经实现。

### 7.3.3 单元测试实现原理

简介
在了解过testing.common后，我们进一步了解testing.T数据结构，以便了解更多单元测试执行的更多细节。

数据结构
源码包src/testing/testing.go:T定义了其数据结构：

```go
type T struct {
    common
    isParallel bool
    context    *testContext // For running tests and subtests.
}
```

其成员简单介绍如下：  
common： 即前面绍的testing.common  
isParallel： 表示当前测试是否需要并发，如果测试中执行了t.Parallel()，则此值为true  
context： 控制测试的并发调度  
因为context直接决定了单元测试的调度，在介绍testing.T支持的方法前，有必要先了解一下context。  

### 7.3.4 性能测试实现原理

#### 简介

跟据前面章节，我们可以快速的写出一个性能测试并执行，最令我感到神奇的是b.N的值，虽然官方资料中说b.N会自动调整以保证可靠的计时，
可还是想了解具体的实现机制。 

本节，我们先分析testing.B数据结构，再看几个典型的成员函数，以期从源码中寻找以下问题的答案： 
- b.N是如何自动调整的？ 
- 内存统计是如何实现的？ 
- SetBytes()其使用场景是什么？

#### 数据结构

源码包src/testing/benchmark.go:B定义了性能测试的数据结构，我们提取其比较重要的一些成员进行分析：  

```go
type B struct {
	common
	importPath       string // import path of the package containing the benchmark
	context          *benchContext
	N                int
	previousN        int           // number of iterations in the previous run
	previousDuration time.Duration // total duration of the previous run
	benchFunc        func(b *B)
	benchTime        benchTimeFlag
	bytes            int64
	missingBytes     bool // one of the subbenchmarks does not have bytes set.
	timerOn          bool
	showAllocResult  bool
	result           BenchmarkResult
	parallelism      int // RunParallel creates parallelism*GOMAXPROCS goroutines
	// The initial states of memStats.Mallocs and memStats.TotalAlloc.
	startAllocs uint64
	startBytes  uint64
	// The net total of this test after being run.
	netAllocs uint64
	netBytes  uint64
	// Extra metrics collected by ReportMetric.
	extra map[string]float64
}
```

其主要成员如下：  
- common： 与testing.T共享的testing.common，管理着日志、状态等；
- N：每个测试中用户代码执行次数
- benchFunc：测试函数
- benchTime：性能测试最少执行时间，默认为1s，可以通过能数-benchtime 2s指定
- bytes：每次迭代处理的字节数
- timerOn：计时启动标志，默认为false，启动计时为true
- startAllocs：测试启动时记录堆中分配的对象数
- startBytes：测试启动时记录堆中分配的字节数
- netAllocs：测试结束后记录堆中新增加的对象数，公式：结束时堆中分配的对象数
- netBytes：测试对事后记录堆中新增加的字节数

**设置处理字节数：B.SetBytes(n int64)**
```go
// SetBytes records the number of bytes processed in a single operation.
// If this is called, the benchmark will report ns/op and MB/s.
func (b *B) SetBytes(n int64) { b.bytes = n }
```
这是一个比较含糊的函数，通过其函数说明很难明白其作用。 

其实它是用来设置单次迭代处理的字节数，一旦设置了这个字节数， 那么输出报告中将会呈现“xxx MB/s”的信息， 用来表示待测函数处理字节的性能。
待测函数每次处理多少字节数只有用户清楚，所以需要用户设置。 

举个例子，待测函数每次执行处理1M数据，如果我们想看待测函数处理数据的性能，那么我们在测试中设置 SetByte(1024 *1024)，
假如待测函数需要执行1s的话，那么结果中将会出现 “1 MB/s”（约等于）的信息。示 例代码如下所示：
