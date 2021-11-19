# Chapter 6 反射

反射是什么？
- 反射是运行时检查自身结构的机智；
- 反射是困惑的源泉

反射特性与interface紧密相关，同时它也与Go类型系统关系密切，这可能都是李杰反射的障碍。

## 6.1 热身测验

#### 1.热身题目  
1). 题目一

如何判断Foo结构体重的两个变量是否相等？是否可以用"=="操作符进行比较？
```go
type Foo struct {
	A int
	B string
	C interface{}
}
```

2). 题目二
如何判断空接口类型中的两个变量是否相等？是否可以用"=="操作符进行比较？
```go
func IsEaqual(a,b interface{}) bool {
	// TODO
}
```
#### 2.参考答案
1). 题目一  
&emsp;&emsp;使用"=="操作符可以比较两个结构体变量，但仅限于结构体成员类型为简单类型，不能包含诸如slice、map等不可比较类型。实际项目中尝尝使用
reflect.DeepEqual()函数来比较两个结构体变量，它支持任意两个结构体变量的比较。  
2). 题目二  
&emsp;&emsp;使用"=="操作符可以比较两个空接口类型变量，但仅限于接口底层类型一致且不包含诸如slice、map等不可比较类型。实际项目尝尝使用
reflect.DeepEqual()函数来比较两个空接口变量，它支持任意接口变量的比较。

## 6.3 反射定律
前面之所以讲类型，是为了引出interface，之所以讲interface是想说interface类型有个(value，type)对， 
而反射就是检查interface的这个(value, type)对的。具体一点说就是Go提供一组方法提取interface的 value，提供另一组方法提取interface的type. 

官方提供了三条定律来说明反射，比较清晰，下面也按照这三定律来总结。 反射包里有两个接口类型要先了解一下. 
- reflect.Type 提供一组接口处理interface的类型，即（value, type）中的type
- reflect.Value 提供一组接口处理interface的值,即(value, type)中的value 

下面会提到反射对象，所谓反射对象即反射包里提供的两种类型的对象。 
reflect.Type 类型对象 
reflect.Value 类型对象

### 2.反射定律

1). 反射第一定律：反射可以将interface类型变量转换成反射对象
注意：反射是针对interface类型变量的，其中 TypeOf() 和 ValueOf() 接受的参数都是 interface{} 类型的，也 即x值是被转成了interface传入的。 
除了 reflect.TypeOf() 和 reflect.ValueOf() ，还有其他很多方法可以操作，本文先不过多介绍，否则一不小心会会 引起困惑。

2). 反射第二定律：反射可以将反射对象还原成interface对象

3). 反射第三定律：反射对象可修改，value值必须是可设置的

## 总结
结合官方博客及本文，至少可以对反射理解个大概，还有很多方法没有涉及。对反射的深入理解，个人觉得还需要继续看的内容： 
- 参考业界，尤其是开源框架中是如何使用反射的 
- 研究反射实现原理，探究其性能优化的手段