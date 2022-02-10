# Chapter 8 异常处理

## 8.1 error

&emsp;&emsp;Ge语言追求极简的设计哲学在error特性上体现得淋澳尽致。

&emsp;&emsp;在Go 1.13之前长达10余年的时间里，标准库对error的支持非常有限，仅有errors.New()
和fmt.Errorf()两个函数用来构造error实例。然而Go语言仅提供了eror的内置接口定义
(type error interface)，这样开发者可以定义任何类型的error，并可以存储任意的内容。

&emsp;&emsp;在Go 1.13之前，已经有很多开源项目试图扩展标准库的error以满足实际项目的需求，比
如pkg/errors,该项目被大量应用于诸如Kubernetes这样的大型项目中。如果读者之前了解过
这类项目，可能会发现Go1.13针对error的优化与这些开源项目有异曲同工之妙。

&emsp;&emsp;Go1.13在保持对原有error兼容的前提下，提供了新的error类型，新的error类型在函数
间传递时可以保存原始的error信息，本节我们称这类error为链式error。

&emsp;&emsp;本节先从基础的error讲起，包括error的基本用法、痛点，接着延伸到Go 1.13的改进，
方便读者了解error的演进脉络。

### 8.1.2 基础error

本节聚焦Go 1.13之前error的设计和使用，包括以下内容：
- error的类型；
- 如何创建error
- 如何自定义error;
- 如何检查error；
- 如何处理error；
- 使用error的痛点；

1.error接口  
2.创建error

### 8.1.3 链式error

fmt.Errorf()讲根据格式动词来动态决定生成wrapError还是errorString。使用%v格式动词生成的error类型仍是errorString(没有实现Unwrap接口)

## 8.2 defer

### 8.2.2 约法三章

1.使用场景   
1). 释放资源  
2). 流程控制  
3). 异常处理

2.行为规则  
1). 规则一：延迟函数的参数在defer语句出现时就已经确定了  
官方给出的一个例子，如下所示：  
```text
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return    
}
```

2). 规则二：延迟函数按后进先出（LIFO）的顺序执行，即先出现的defer最后执行  

3). 规则三：延迟函数可能操作主函数的具名返回值  
定义defer的函数（下称主函数）可能有返回值，**返回值可能有名字（具名返回值）**，也可能没有名字（匿名返回值）延迟函数可能会影响返回值。

若要理解延迟函数是如何影响主函数返回值的，只要明白函数是如何返回的就足够了。  

1). 函数返回过程  
有一个事实必须要了解，关键字return不是一个原子操作，实际上return只代表汇编指令ret，即跳转程序执行。比如
ret, 即跳转程序执行。比如语句return i，实际上分两步执行，即先将i值存人栈中作为返回值，然后执行跳转，
而defer的执行时机正是在跳转前，所以说dr执行时还是有机会操作返回值的。

举个实际的例子来说明这个过程：  
```text
func deferFuncReturn()(result int){
    i:=1
    defer func()(
        result++
    }()
    return i
}
```
该函数的return语句可以拆分成下面两行：  
```text
result = i  
return  
```
而延迟函数的执行正是在return之前，即加入defer后的执行过程如下：
```text
result = i  
result++  
return  
```
所以上面的函数实际返回i++值。  
主函数有不同的返回方式，包括匿名返回值和具名返回值，但万变不离其宗，只要把return语句拆开都可以很好理解，下面分别举例说明。  

(2)主函数拥有匿名返回值，返回字面值。  
一个主函数拥有一个匿名返回值，返回时使用字面值，比如返回1、2、Hello这样的值，这种情况下defer语句是无法操作返回值的。

一个返回字面值的函数如下：
```text
func foo() int {
    var i int
    
    defer func() {
        i++
    }
    
    return 1
}
```
上面的ruturn语句直接把1写人栈中作为返回值，延迟函数无法操作该返回值，所以就无法影响返回值。

(3)主函数拥有匿名返回值，返回变量。

一个主函数拥有一个匿名返回值，返回本地或全局变量，这种情况下defer语句可以引用返回值，但不会改变返回值。

一个返回本地变量的函数如下：
```text
func foo() int {
    var i int
    
    defer func(){
        i++
    }()
    
    return i
}
```
上面的函数返回一个局部变量，同时defer函数也会操作这个局部变量。对于匿名返回值来说，可以假定仍然有一个变量存储返回值，
假定返回值变量为anony, 上面的返回语句可以拆分为以下过程：
```text
anony = i
i++
return
```
由于i是整型值，会将值复制给anony,所以在defer语句中修改i值，不会对函数返回值造成影响。

(4)主函数拥有具名返回值。  

主函数声明语句中带名字的返回值会被初始化为一个局部变量，函数内部可以像使用局部变量一样使用该返回值。如果defer语句操作该
返回值，则可能改变返回结果。

一个影响函数返回值的例子：  
```text
func foo() (ret int) {
    defer func() {
        ret++
    }()
    
    return 0
}
```
上面的函数拆解出来如下所示：  
```text
ret = 0 
ret++
return
```
函数真正返回前，在defer中对返回值做了+1操作，所以最终返回1

### 8.2.3 实现原理

1.数据结构  
src/runtime/runtime2.go:_defer  

2.defer的创建和执行  

源码包中src/ime/panic.go定义了以下两个方法，分别用于创建defer和执行defer。  
deferproc():   用于将defer函数处理成defer实例，并存人goroutine链表中；  
deferreturn(): 用于将defer从goroutine链表中取出并执行。  

可以简单地这么理解，编译器在编译阶段把defer语句替换成了函数deferproc(),在函数
return前插人了函数deferreturnO。在运行时，每执行一次deferproc)都创建一个运行时_defer
实例并存储，函数返回前执行deferreturn()依次取出_defer实例并执行。

3.小结
- defer定义的延迟函数参数在defer语句出现时就已经确定了；
- defer定义的顺序与实际的执行顺序相反；
- return不是原子操作，执行过程是：保存返回值（若有）->执行defer（若有）->执行ret跳转；
- 申请资源后立即使用defer关闭资源是一个好习惯。

### 8.2.4 性能优化

defer 的实现机制在Go的演进过程中不断被优化，根据其实现机制的不同，可以把defer分为三种类型。
- heap-allocated:存储在堆上的defer, Go 1.13之前只有这种类型。
- stack-allocated:存储在栈上的defer,Go 1.13引入这种类型。
- open-coded:开放编码类型的defer, Go 1.14引入这种类型。

前面章节介绍的defer实际上是第一种类型，即存储在堆上的defer,但在Go 1.13及Go
1.14引入新的类型后，原类型也没有被废弃，而是多种类型共存，编译器会决策使用何种类型
优先使用open-coded(以下简称开放编码defer),stack-allocated(以下简称栈defer)次之，在
有些编译器无法决定的场景或受限的场景中则会使用heap-allocated(以下简称堆defer)。
defer的编译和执行需要编译器和Go runtime的配合，本节所指的“运行时”往往是指Go的runtime组件。

1.堆defer  
&emsp;&emsp;与其后出现的两种defer类型相比较，堆defer的痛点主要在于频繁的堆内存分配以及释放，性能稍差  
2.栈defer  
&emsp;&emsp;编译器会尽可能地把defer语句编译成栈类型，但由于栈空间有限，并不能把所有的defer都存储在栈中，所以还需要保留堆defer。  
3.开放编码defer  
1). 禁用编译器优化  
2). 循环语句  
3). 限制defer数量

## 8.3 panic

### 8.3.2 工作机制

1. panic()函数
2. 工作流程  
    
3. 在panic的执行过程中有几个要点要明确：
    - panic会递归执行协程中所有的defer，与函数正常退出时的执行顺序一致；
    - panic不会处理其他协程中的defer；
    - 当前协程中的defer处理完成后，触发程序退出。
4. 小结

### 8.3.3 源码剖析

1.panic的真身  
2.数据结构  
3.gopanic分析  
&emsp;1). 没有defer函数  
&emsp;2). defer函数处理  
&emsp;3). 嵌套defer  

## 8.4 recover

内置函数recover()用于消除panic并使程序恢复正常，看起来很简单，但下面有几个问题是否困扰到你：
- recover()的返回值是什么？
- 执行recover()之后程序将从哪里继续运行？
- recover()为什么一定要在defer函数中使用？

### 8.4.2 工作机制

1.recover()函数

recover()函数的返回值就是panic()函数的参数，当程序产生panic时，recover()函数就可用
于消除panic.同时返回panic()函数的参数，如果程序没有发生panic,则recover()函数返回nil.

如果panic()函数参数为nil,那么仍然是一个有效的panic,此时recover()函数仍然可以捕
获panic，但返回值为nil。如以下代码所示，由于err为nil,所以字符串A得不到打印，
panic可以被消除：
```text
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	panic(nil)
	fmt.Println("B")
```
2.工作流程  

3.小结  
关于recover()函数，有几个要点要明确：
- recover()函数调用必须要位于defer函数中，且不能出现在另一个嵌套函数中；
- recover()函数成功处理异常后，无法再次回到本函数发生panic的位置继续执行；
- recover()函数可以消除本函数产生或收到的panic，上游函数感知不到panic的发生。

当函数中发生panic并用recover()函数恢复后，当前函数仍然会继续返回，对于匿名返回
值，函数将返回相应类型的零值，对于具名返回值，函数将返回当前已经存在的值。

### 8.3.4 源码剖析

1.recover的真身  

2.gorecover()  
1). 恢复逻辑  
2). 生效条件  
3). 设计思路  

3.小结  
再次总结一下recover()函数生效的条件：
- 程序中必须真正产生了panic
- recover()函数在defer函数中；
- recover()函数被defer函数直接调用；
- 一般除了 HTTP、RPC 服务，其他的场景都不需要 recover 的，panic 说明程序运行的问题没有被考虑到，
找到办法解决就可以了，往往是好事