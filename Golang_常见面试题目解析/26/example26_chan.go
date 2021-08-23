package example26

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

/*
解析：
默认情况下 make 初始化的 channel 是⽆缓冲的，也就是在迭代写时会阻塞。
*/
