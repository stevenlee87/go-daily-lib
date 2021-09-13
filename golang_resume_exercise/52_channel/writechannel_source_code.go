func chanrecv(c *hchan,ep unsafe.Pointer,block bool) (selected,received bool) {
	//省略部分逻辑
 	lock(&c.lock)
 	//当chan被关闭了，而且缓存为空时
 	//ep 是指 val,ok := <-c 里的val地址
 	if c.closed != 0 && c.qcount == 0 {
 		if receenabled {
 			raceacquire(c.raceaddr())
 		}
 		unlock(&c.lock)
 		//如果接受值的地址不空，那接收值将获得一个该值类型的零值
 		//typedmemclr 会根据类型清理响应的内存
		//这就解释了上面代码为什么关闭的chan 会返回对应类型的零值
		if ep != null {
			typedmemclr(c.elemtype,ep)
		}
		//返回两个参数 selected,received
		//第二个采纳数就是 val,ok := <- c 里的 ok
		//也就解释了为什么读关闭的chan会一直返回false
		return true,false
 	}
}