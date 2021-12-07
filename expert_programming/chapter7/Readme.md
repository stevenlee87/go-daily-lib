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
假如待测函数需要执行1s的话，那么结果中将会出现 “1 MB/s”（约等于）的信息。示例代码如下所示：
```go
func BenchmarkSetBytes(b *testing.B) {
    b.SetBytes(1024 * 1024)
    for i := 0; i < b.N; i++ {
        time.Sleep(1 * time.Second) // 模拟待测函数
    }
}
```

打印结果：
```go
goos: darwin
goarch: amd64
pkg: github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkSetBytes
BenchmarkSetBytes-8   	       1	1004217211 ns/op	   1.04 MB/s
PASS
```

可以看到测试执行了一次，花费时间约1S，数据处理能力约为1MB/s。

B.N是如何调整的？  
B.launch()方法里最终决定B.N的值。我们看下伪代码：  
```go
func (b *B) launch() { // 此方法自动测算执行次数，但调用前必须调用run1以便自动计算次数
    d := b.benchTime
    for n := 1; !b.failed && b.duration < d && n < 1e9; { // 最少执行b.benchTime（默认为1s）时间，最多执行1e9次
        last := n
        n = int(d.Nanoseconds()) // 预测接下来要执行多少次，b.benchTime/每个操作耗时
        if nsop := b.nsPerOp(); nsop != 0 {
            n /= int(nsop)
        }
        n = max(min(n+n/5, 100*last), last+1) // 避免增长较快，先增长20%，至少增长1次
        n = roundUp(n) // 下次迭代次数向上取整到10的指数，方便阅读
        b.runN(n)
    }
}
```

不考虑程序出错，而且用户没有主动停止测试的场景下，每个性能测试至少要执行b.benchTime长的秒数，默认为1s。
先执行一遍的意义在于看用户代码执行一次要花费多长时间，如果时间较短，那么b.N值要足够大才可以测得更精确，如果时间较长，b.N值相应的会减少，
否则会影响测试效率。

最终的b.N会被定格在某个10的指数级，是为了方便阅读测试报告。

**内存是如何统计的？**

我们知道在测试开始时，会把当前内存值记入到b.startAllocs和b.startBytes中，测试结束时，会用最终内存值与开始时的内存值相减，
得到净增加的内存值，并记入到b.netAllocs和b.netBytes中。

每个测试结束，会把结果保存到BenchmarkResult对象里，该对象里保存了输出报告所必需的统计信息：  
```go
// BenchmarkResult contains the results of a benchmark run.
type BenchmarkResult struct {
	N         int           // The number of iterations.
	T         time.Duration // The total time taken.
	Bytes     int64         // Bytes processed in one iteration.
	MemAllocs uint64        // The total number of memory allocations.
	MemBytes  uint64        // The total number of bytes allocated.

	// Extra records additional metrics reported by ReportMetric.
	Extra map[string]float64
}
```

其中MemAllocs和MemBytes分别对应b.netAllocs和b.netBytes。

那么最终统计时只需要把净增加值除以b.N即可得到每次新增多少内存了。

每个操作内存对象新增值：
```go
// AllocsPerOp returns the "allocs/op" metric,
// which is calculated as r.MemAllocs / r.N.
func (r BenchmarkResult) AllocsPerOp() int64 {
	if v, ok := r.Extra["allocs/op"]; ok {
		return int64(v)
	}
	if r.N <= 0 {
		return 0
	}
	return int64(r.MemAllocs) / int64(r.N)
}
```

每个操作内存字节数新增值：  
```go
// AllocedBytesPerOp returns the "B/op" metric,
// which is calculated as r.MemBytes / r.N.
func (r BenchmarkResult) AllocedBytesPerOp() int64 {
	if v, ok := r.Extra["B/op"]; ok {
		return int64(v)
	}
	if r.N <= 0 {
		return 0
	}
	return int64(r.MemBytes) / int64(r.N)
}
```

### 7.3.5 示例测试实现原理

简介
示例测试相对于单元测试和性能测试来说，其实现机制比较简单。它没有复杂的数据结构，也不需要额外的流程控制，其核心工作原理在于收集测试过程中的打印日志，然后与期望字符串做比较，最后得出是否一致的报告。

testing/example.go  
**数据结构**

```go
type InternalExample struct {
    Name      string    // 测试名称
    F         func()   // 测试函数
    Output    string    // 期望字符串
    Unordered bool      // 输出是否是无序的
}
```

比如，示例测试如下：  
```go
// 检测乱序输出
func ExamplePrintNames() {
    gotest.PrintNames()
    // Unordered output:
    // Jim
    // Bob
    // Tom
    // Sue
}
```

该示例测试经过编译后，产生的数据结构成员如下：

- InternalExample.Name = "ExamplePrintNames";
- InternalExample.F = ExamplePrintNames()
- InternalExample.Output = "Jim\n Bob\n Tom\n Sue\n"
- InternalExample.Unordered = true;
其中Output是包含换行符的字符串。

### 7.3.6 Main测试的实现原理

**简介**  
每一种测试（单元测试、性能测试或示例测试），都有一个数据类型与其对应。

单元测试：InternalTest  
性能测试：InternalBenchmark  
示例测试：InternalExample  
测试编译阶段，每个测试都会被放到指定类型的切片中，测试执行时，这些测试将会被放到testing.M数据结构中进行调度。
而testing.M即是MainTest对应的数据结构。

**数据结构**  
源码src\testing/testing.go:M定义了testing.M的数据结构：
```go
// M is a type passed to a TestMain function to run the actual tests.
type M struct {
	deps       testDeps
	tests      []InternalTest
	benchmarks []InternalBenchmark
	examples   []InternalExample

	timer     *time.Timer
	afterOnce sync.Once

	numRun int

	// value to pass to os.Exit, the outer test func main
	// harness calls os.Exit with this code. See #34129.
	exitCode int
}
```

单元测试、性能测试和示例测试在经过编译后都会被存放到一个testing.M数据结构中，在测试执行时该数据结构将传递给TestMain()，
真正执行测试的是testing.M的Run()方法，这个后面我们会继续分析。

timer用于指定测试的超时时间，可以通过参数timeout <n>指定，当测试执行超时后将会立即结束并判定为失败。

**执行测试**

TestMain()函数通常会有一个m.Run()方法，该方法会执行单元测试、性能测试和示例测试，如果用户实现了TestMain()但没有调用m.Run()的话，
那么什么测试都不会被执行。

m.Run()不仅会执行测试，还会做一些初始化工作，比如解析参数、起动定时器、跟据参数指示创建一系列的文件等。

m.Run()使用三个独立的方法来执行三种测试：

- 单元测试：runTests(m.deps.MatchString, m.tests)
- 性能测试：runExamples(m.deps.MatchString, m.examples)
- 示例测试：runBenchmarks(m.deps.ImportPath(), m.deps.MatchString, m.benchmarks)
其中m.deps里存放了测试匹配相关的内容，暂时先不用关注。

### 7.3.7 go test的工作机制

**前言**  
前面的章节我们分析了每种测试的数据结构及其实现原理，本节我们看一下go test的执行机制。

Go 有多个命令行工具，go test只是其中一个。go test命令的函数入口在src\cmd\go\internal\test\test.go:runTest()，
这个函数就是go test的大脑。

**两种运行模式**  
go test运行时，跟据是否指定package分为两种模式，即本地目录模式和包列表模式。

**本地目录模式**  
当执行测试并没有指定package时，即以本地目录模式运行，例如使用"go test"或者"go test -v"来启动测试。

本地目录模式下，go test编译当前目录的源码文件和测试文件，并生成一个二进制文件，最后执行并打印结果。

**包列表模式**  
当执行测试并显式指定package时，即以包列表模式运行，例如使用"go test math"来启动测试。

包列表模式下，go test为每个包生成一个测试二进制文件，并分别执行它。 包列表模式是在Go 1.10版本才引入的，它会把每个包的测试结果写入到本地临时
文件中做为缓存，下次执行时会直接从缓存中读取测试结果，以便节省测试时间。

**缓存机制**
当满足一定的条件，测试的缓存是自动启用的，也可以显式的关闭缓存。

**测试结果缓存**
如果一次测试中，其参数全部来自"可缓存参数"集合，那么本次测试结果将被缓存。

**可缓存参数集合如下**：

- -cpu
- -list
- -parallel
- -run
- -short
- -v  
需要注意的是，测试参数必须全部来自这个集合，其结果才会被缓存，没有参数或包含任一此集合之外的参数，结果都不会缓存。

**使用缓存结果**  
如果满足条件，测试不会真正执行，而是从缓存中取出结果并呈现，结果中会有"cached"字样，表示来自缓存。

使用缓存结果也需要满足一定的条件：

- 本次测试的二进制及测试参数与之前的一次完全一致；
- 本次测试的源文件及环境变量与之前的一次完全一致；
- 之前的一次测试结果是成功的；
- 本次测试运行模式是列表模式；

下面演示一个使用缓存的例子：
```go
go test .
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest  1.273s
go test .
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest  (cached)
```
前后两次执行测试，参数没变，源文件也没变化，第二次执行时会自动从缓存中获取结果，结果中“cached”即表示结果从缓存中获取。

**禁用缓存**  
测试时使用一个不在“可缓存参数”集合中的参数，就不会使用缓存，比较常用的方法是指定一个参数“-count=1”。

下面演示一个禁用缓存的例子：
```go
go test .
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest  1.273s
go test .
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest  (cached)
go test . -count=1
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest  0.565s
```

第三次执行使用了参数"-count=1"，所以执行时不会从缓存中获取结果。

## 7.4 扩展阅读

go test非常容易上手，但并不代表其功能单一，它提供了丰富的参数接口以便满足各种测试场景。

本节，我们主要介绍一些常用的参数，通过前面实现原理的学习和本节的示例，希望读者可以准确掌握其用法，以便在工作中提供便利。

### 7.4.1 测试参数

**前言**  
go test有非常丰富的参数，一些参数用于控制测试的编译，另一些参数控制测试的执行。

有关测试覆盖率、vet和pprof相关的参数先略过，我们在讨论相关内容时再详细介绍。

**控制编译的参数**  
**-args**  
指示go test把-args后面的参数带到测试中去。具体的测试函数会跟据此参数来控制测试流程。

-args后面可以附带多个参数，所有参数都将以字符串形式传入，每个参数做为一个string，并存放到字符串切片中。
```go
// TestArgs 用于演示如何解析-args参数
func TestArgs(t *testing.T) {
    if !flag.Parsed() {
        flag.Parse()
    }

    argList := flag.Args() // flag.Args() 返回 -args 后面的所有参数，以切片表示，每个元素代表一个参数
    for _, arg := range argList {
        if arg == "cloud" {
            t.Log("Running in cloud.")
        }else {
            t.Log("Running in other mode.")
        }
    }
}
```

执行测试时带入参数：
```go
go test . -run TestArgs -v -args "cloud"
=== RUN   TestArgs
    testargs_test.go:16: Running in cloud.
--- PASS: TestArgs (0.00s)
PASS
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest  3.054s
```
通过参数-args指定传递给测试的参数。

**-json**  
-json 参数用于指示go test将结果输出转换成json格式，以方便自动化测试解析使用。

示例如下：
```go
go test -run TestAdd -json
{"Time":"2021-12-07T15:22:07.760538+08:00","Action":"run","Package":"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest","Test":"TestAdd"}
{"Time":"2021-12-07T15:22:07.760926+08:00","Action":"output","Package":"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest","Test":"TestAdd","Output":"=== RUN   TestAdd\n"}
{"Time":"2021-12-07T15:22:07.760973+08:00","Action":"output","Package":"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest","Test":"TestAdd","Output":"--- PASS: TestAdd (0.00s)\n"}
{"Time":"2021-12-07T15:22:07.760992+08:00","Action":"pass","Package":"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest","Test":"TestAdd","Elapsed":0}
{"Time":"2021-12-07T15:22:07.761019+08:00","Action":"output","Package":"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest","Output":"PASS\n"}
{"Time":"2021-12-07T15:22:07.761075+08:00","Action":"output","Package":"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest","Output":"ok  \tgithub.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest\t1.034s\n"}
{"Time":"2021-12-07T15:22:07.761104+08:00","Action":"pass","Package":"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest","Elapsed":1.035}
```

**-o**  
-o 参数指定生成的二进制可执行程序，并执行测试，测试结束不会删除该程序。

没有此参数时，go test生成的二进制可执行程序存放到临时目录，执行结束便删除。

示例如下：
```go
go test -run TestAdd -o TestAdd
PASS
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest  3.078s
```

本例中，使用-o 参数指定生成二进制文件"TestAdd"并存放到当前目录，测试执行结束后，仍然可以直接执行该二进制程序。

控制测试的参数  
-bench regexp  
go test默认不执行性能测试，使用-bench参数才可以运行，而且只运行性能测试函数。

其中正则表达式用于筛选所要执行的性能测试。如果要执行所有的性能测试，使用参数"-bench ."或"-bench=."。

此处的正则表达式不是严格意义上的正则，而是种包含关系。

比如有如下三个性能测试：

- func BenchmarkMakeSliceWithoutAlloc(b *testing.B)
- func BenchmarkMakeSliceWithPreAlloc(b *testing.B)
- func BenchmarkSetBytes(b *testing.B)
使用参数“-bench=Slice”，那么前两个测试因为都包含"Slice"，所以都会被执行，第三个测试则不会执行。

对于包含子测试的场景下，匹配是按层匹配的。举一个包含子测试的例子：
```go
func BenchmarkSub(b *testing.B) {
    b.Run("A=1", benchSub1)
    b.Run("A=2", benchSub2)
    b.Run("B=1", benchSub3)
}
```

测试函数命名规则中，子测试的名字需要以父测试名字做为前缀并以"/"连接，上面的例子实际上是包含4个测试：

- Sub
- Sub/A=1
- Sub/A=2
- Sub/B=1  
如果想执行三个子测试，那么使用参数“-bench Sub”。如果只想执行“Sub/A=1”，则使用参数"-bench Sub/A=1"。如果想执行"Sub/A=1"和“Sub/A=2”，
则使用参数"-bench Sub/A="。

-benchtime s  
-benchtime指定每个性能测试的执行时间，如果不指定，则使用默认时间1s。

例如，执定每个性能测试执行2s，则参数为："go test -bench Sub/A=1 -benchtime 2s"。

-cpu   
参数提供一个CPU个数的列表，提供此列表后，那么测试将按照这个列表指定的CPU数设置GOMAXPROCS并分别测试。

比如“-cpu 1,2”，那么每个测试将执行两次，一次是用1个CPU执行，一次是用2个CPU执行。 例如，使用命令"go test -bench Sub/A=1 -cpu 1,2,3,4" 执行测试：

```go
go test -bench Sub/A=1 -cpu 1,2,3,4
goos: darwin
goarch: amd64
pkg: github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkSub/A=1                    2317            497982 ns/op
BenchmarkSub/A=1-2                  2617            437995 ns/op
BenchmarkSub/A=1-3                  2557            457854 ns/op
BenchmarkSub/A=1-4                  2491            489951 ns/op
PASS
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start 5.239s

```
测试结果中测试名后面的-2、-3、-4分别代表执行时GOMAXPROCS的数值。 如果GOMAXPROCS为1，则不显示。