## Chapter 5 并发控制

本章主要介绍GO语言开发过程中经常使用的并发控制手段。

我们考虑这么一种场景，协程A执行过程中需要创建子协程A1、A2、A3…An，协程A创建完子协程后就等待子协程退

出。针对这种场景，GO提供了三种解决方案：

Channel: 使用channel控制子协程

WaitGroup : 使用信号量机制控制子协程

Context: 使用上下文控制子协程

三种方案各有优劣，比如Channel优点是实现简单，清晰易懂，WaitGroup优点是子协程个数动态可调整，Context

优点是对子协程派生出来的孙子协程的控制。缺点是相对而言的，要结合实例应用场景进行选择。 

## Context
Golang context是Golang应用开发常用的并发控制技术，它与WaitGroup最大的不同点是context对于派生 goroutine有更强的控制力，
它可以控制多级的goroutine。 context翻译成中文是”上下文”，即它可以控制一组呈树状结构的goroutine，每个goroutine拥有相同的上下文。
典型的使用场景如下图所示：
![avatar](https://github.com/stevenlee87/go-daily-lib/blob/main/expert_programming/chapter5/5.3_context/test.go)

## 总结
- Context仅仅是一个接口定义，跟据实现的不同，可以衍生出不同的context类型；
- cancelCtx实现了Context接口，通过WithCancel()创建cancelCtx实例；
- timerCtx实现了Context接口，通过WithDeadline()和WithTimeout()创建timerCtx实例；
- valueCtx实现了Context接口，通过WithValue()创建valueCtx实例；
- 三种context实例可互为父节点，从而可以组合成不同的应用形式； 