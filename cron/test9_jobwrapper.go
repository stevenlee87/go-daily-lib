package main

/*
cron对象创建使用了选项模式，我们前面已经介绍了 3 个选项：

WithLocation：指定时区；
WithParser：使用自定义的解析器；
WithSeconds：让时间格式支持秒，实际上内部调用了WithParser。

cron还提供了另外两种选项：
WithLogger：自定义Logger；
WithChain：Job 包装器。

cron内置了 3 个用得比较多的JobWrapper：
Recover：捕获内部Job产生的 panic；
DelayIfStillRunning：触发时，如果上一次任务还未执行完成（耗时太长），则等待上一次任务完成之后再执行；
SkipIfStillRunning： 触发时，如果上一次任务还未执行完成，则跳过此次执行。
*/

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type panicJob struct {
	count int
}

func (p *panicJob) Run() {
	p.count++
	if p.count == 1 {
		panic("oooooooooooooops!!!")
	}

	fmt.Println("hello world")
}

func main() {
	c := cron.New()
	c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&panicJob{}))
	c.Start()

	time.Sleep(5 * time.Second)
}
