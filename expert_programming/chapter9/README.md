# Chapter 9 定时器

定时器在Go语言应用中使用非常广泛，准确掌握其用法和实现原理至关重要。  
Go提供了两种定时器，此处分为一次性定时器、周期性定时器。  
- 一次性定时器：定时器只计时一次，结束便停止； 
- 周期性定时器：定时器周期性进行计时，除非主动停止，否则将永久运行；  

本章会快速介绍这两种定时器的基本用法，重点介绍其内部实现原理，最后再给出一个案例揭示使用定时器的风险。

## 9.1 一次性定时器（Timer）

### 9.1.1 快速开始

1. 简介  

2. 使用场景  
1). 设定超时时间  
2). 延迟执行某个方法

3. Timer对外接口  
1). 创建定时器  
2). 停止定时器  
3). 重置定时器

4. 简单接口  
1). After()  
2). AfterFunc()

5. 小结  
Timer内容总结如下：
   - time.NewTimer(d)创建一个Timer; 
   - timer.Stop()停掉当前Timer; 
   - timer.Reset(d)重置当前Timer;

### 9.1.2 实现原理

很多人想当然的以为，启动一个Timer意味着启动了一个协程，这个协程会等待Timer到期，
然后向Timer的管道中 发送当前时间。 

实际上，每个Go应用程序都有一个协程专门负责管理所有的Timer，这个协程负责监控Timer是否过期，
过期后执行一个预定义的动作， 这个动作对于Timer而言就是发送当前时间到管道中。

1. 数据结构  
1). Timer  
源码包 src/time/sleep.go:Timer 定义了其数据结构：  
```text
type Timer struct {
   C <-chan Time
   r runtimeTimer
}
```  

2). runtimeTimer  

2. 实现原理  

一个进程中的多个Timer都由底层的一个协程来管理，为了描述方便我们把这个协程称为系统协程。 

我们想在后面的章节中单独介绍系统协程工作机制，本节，我们先简单介绍其工作过程。 

系统协程把runtimeTimer存放在数组中，并按照 when 字段对所有的runtimeTimer进行堆排序，
定时器触发时执行runtimeTimer中的预定义函数 f ，即完成了一次定时任务。  

1). 创建Timer  
2). 停止Timer  
3). 重置Timer

3. 小结  
- NewTimer()创建一个新的Timer交给系统协程监控； 
- Stop()通知系统协程删除指定的Timer; 
- Reset()通知系统协程删除指定的Timer并再添加一个新的Timer；

## 9.2 周期性定时器(Ticker)

### 9.2.1 快速开始

1. 简介

Ticker是周期性定时器，即周期性的触发一个事件，通过Ticker本身提供的管道将事件传递出去。 

Ticker的数据结构与Timer完全一致： 

```text
type Ticker struct { 
    C <-chan Time
    r runtimeTimer
} 
```

Ticker对外仅暴露一个channel，指定的时间到来时就往该channel中写入系统时间，也即一个事件。 

在创建Ticker时会指定一个时间，作为事件触发的周期。这也是Ticker与Timer的最主要的区别。 

另外，ticker的英文原意是钟表的”滴哒”声，钟表周期性的产生”滴哒”声，也即周期性的产生事件。

2. 使用场景  
1). 简单定时任务  
2). 简单定时任务

3. Ticker对外接口

4. 简单接口

5. 错误示例

6. 小结