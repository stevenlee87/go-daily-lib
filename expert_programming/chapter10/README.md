# Chapter 10 语法糖

## 名字由来

语法糖（Syntactic sugar）的概念是由英国计算机科学家 Peter J. Landin 提出的，
用于表示编程语言中的某种 类型的语法，这些语法不会影响功能，但使用起来却很方便。 

语法糖，也称糖语法，这些语法不仅不会影响功能，编译后的结果跟不使用语法糖也一样。 

语法糖，有可能让代码编写变得简单，也有可能让代码可读性更高，也有可能让代码出问题。
为了避免陷阱才是这个章节的重点。

## Go语言语法糖

最常用的语法糖莫过于赋值符 := ，其次，表示函数变参的 ... 。 

接下来，我们会介绍这两种语法糖的用法，更重要的是结合实际的经历跟大家分享其中的陷阱。

## 10.1 简短变量声名符

想要声明变量，可以使用关键字 var 或者直接使用简短变量声明（ := ）。后者使用更频繁一些，尤其是在接收函数返回值场景中，
你不必使用 var 声明一个变量再变量接收函数返回值，使用 := 可以一步到位。 

本节我们讨论 := 的一些容易被忽视的规则，以避免一些陷阱。

### 10.1.2 规则  

虽然简短变量声明这个语法糖用起来很方便，但有时也会给你一个意外也可能带你掉入陷阱。 

我曾因滥用这个 := 语法糖，发生过一次故障，所以才认真研究了一下它的原理和规则，大家可以做为参考。

1. 规则一： 多变量赋值可能会重新声明

我们知道使用 := 一次可以声明多个变量，像下面这样：
```text
field1, offset := nextField(str, 0)
```
上面代码定义了两个变量，并用函数返回值进行赋值。 

如果这两个变量中的一个再次出现在 := 左侧就会重新声明。像下面这样：

```text
 field1, offset := nextField(str, 0)
 field2, offset := nextField(str, offset)
```
offset被重新声明。 
重新声明并没有什么问题，它并没有引入新的变量，只是把变量的值改变了，但要明白， 这是Go提供的一个语法糖。

- 当 := 左侧存在新变量时（如field2），那么已声明的变量（如offset）则会被重新声明，不会有其他额外副作用。 
- 当 := 左侧没有新变量是不允许的，编译会提示 no new variable on left side of := 。 

我们所说的重新声明不会引入问题要满足一个前提，变量声明要在同一个作用域中出现。如果出现在不同的作用域，
那很可能就创建了新的同名变量，同一函数不同作用域的同名变量往往不是预期做法，很容易引入缺陷。关于作用域
的这个问题，我们在本节后面介绍。

2. 规则二：不能用于函数外部

简短变量场景只能用于函数中，使用 := 来声明和初始化全局变量是行不通的。 比如，像下面这样：
```text
package main

import "fmt"

rule := "Short"
const ttt = 1

func main() {
    fmt.Println("test")
}

/*
syntax error: non-declaration statement outside function body
*/
```

这里的编译错误提示 syntax error: non-declaration statement outside function body ，
表示非声明语句不能出现在函数外部。可以理解成 := 实际上会拆分成两个语句，即声明和赋值。
赋值语句不能出现在函数外部的。

3.变量作用域问题

几乎所有的工程师都了解变量作用域，但是由于 := 使用过于频繁的话，还是有可能掉进陷阱里。

下面代码源自真实项目，但为了描述方便，也为了避免信息安全风险，简化如下：
```text
func Redeclare() {
    field, err := nextField() // 1号err

    if field == 1 {
        field, err := nextField()    // 2号err
        newField, err := nextField() // 3号err
        ...
    }
    ...
}
```
注意上面声明的三个err变量。
- 2号err与1号err不属于同一个作用域，:= 声明了新的变量，所以2号err与1号 err属于两个变量。
- 2号err与3号err属于同一个作用域， := 重新声明了err但没创建新的变量，所以2号err与3号err是同一个变量。 

如果误把2号err与1号err混淆，就很容易产生意想不到的错误。

## 10.2 可变参函数

可变参函数是指函数的某个参数可有可无，即这个参数个数可以是0个或多个。声明可变参数函数的方式是在参数类型 前加上 ... 前缀。 

比如 fmt 包中的 Println :
```text
func Println(a ...interface{})
```

1. 函数特征

我们先写一个可变参函数：
```text
func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.\n\n")
		return
	}

	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}
```

Greeting 函数负责给指定的人打招呼，其参数 who 为可变参数。 

这个函数几乎把可变参函数的特征全部表现出来了：
- 可变参数必须在函数参数列表的尾部，即最后一个（如放前面会引起编译时歧义）；
- 可变参数在函数内部是作为切片来解析的；
- 可变参数可以不填，不填时函数内部当成 nil 切片处理；
- 可变参数必须是相同类型的（如果需要是不同类型的可以定义为interface{}类型）；

2. 使用举例

我们使用 testing 包中的Example函数来说明上面 Greeting 函数（函数位于sugar包中）用法。

1). 不传值

调用可变参函数时，可变参部分是可以不传值的，例如：
```text
func ExampleGreetingWithoutParameter() { 
    sugar.Greeting("nobody")
    // OutPut: 
    // Nobody to say hi
}
```
这里没有传递第二个参数。可变参数不传递的话，默认为nil。

2). 传递多个参数

调用可变参函数时，可变参数部分可以传递多个值，例如：
```text
func ExampleGreetingWithParameter() { 
    sugar.Greeting("hello:", "Joe", "Anna", "Eileen")
    // OutPut: 
    // hello: Joe 
    // hello: Anna 
    // hello: Eileen 
}
```
可变参数可以有多个。多个参数将会生成一个切片传入，函数内部按照切片来处理。

3). 传递切片

调用可变参函数时，可变参数部分可以直接传递一个切片。参数部分需要使用 slice... 来表示切片。例如：
```text
func ExampleGreetingWithSlice() { 
    guest := []string{"Joe", "Anna", "Eileen"} 
    sugar.Greeting("hello:", guest...) 
    // OutPut: 
    // hello: Joe 
    // hello: Anna 
    // hello: Eileen 
}
```
此时需要注意的一点是，切片传入时不会生成新的切片，也就是说函数内部使用的切片与传入的切片共享相同的存储空间。
说得再直白一点就是，如果函数内部修改了切片，可能会影响外部调用的函数。

3. 小结

- 可变参数必须要位于函数列表尾部； 
- 可变参数是被当作切片来处理的； 
- 函数调用时，可变参数可以不填； 
- 函数调用时，可变参数可以填入切片；