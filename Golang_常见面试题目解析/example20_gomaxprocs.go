package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
其中Go1.13.8和Go1.14.6，在实现上和早期版本有一点不同，每增加一个子协程就把其对应的函数地址存放到”_p_.runnext“，
而把”_p_.runnext“原来的地址（即上一个子协程对应的函数地址）移动到队列”_p_.runq“里面，
这样当执行到wg.Wait()时，”_p_.runnext“存放的就是最后一个子协程对应的函数地址（即输出B: ９的那个子协程）。

当开始执行子协程对应的函数时，首先执行”_p_.runnext“对应的函数，然后按先进先出的顺序执行队列”_p_.runq“里的函数。
所以这就解释了为什么总是B：9打在第一个，而后面打印的则是进入队列的顺序。
*/
