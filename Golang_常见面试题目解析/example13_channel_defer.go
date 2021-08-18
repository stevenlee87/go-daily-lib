package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		defer p.deferError()
		m := msg.(int)
		fmt.Println("msg: ", m)
	}
}

func (p *Project) run(msgchan chan interface{}) {
	for {
		//defer p.deferError() // defer p.deferError() 需要在协程开始出调⽤，否则⽆法捕获  panic
		go p.exec(msgchan)

		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			time.Sleep(time.Second)
		}
	}()
	//time.Sleep(time.Second * 100000000000000) // constant 100000000000000000000000 overflows time.Duration
	time.Sleep(time.Second * 10)
}

func main() {
	p := new(Project)
	p.Main()
}
