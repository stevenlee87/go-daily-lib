package main

// 8、下⾯的代码有什么问题?
type UserAges struct {
	ages map[string]int
	//sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	//ua.Lock()
	//defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {

		return age
	}
	return -1
}

func main() {
	ua := UserAges{map[string]int{"lee": 18}}
	ua.Get("lee")
	ua.Add("zhao", 20)

}

/*
解析：
在执⾏ Get⽅法时可能被panic。
虽然有使⽤sync.Mutex做写锁，但是map是并发读写不安全的。map属于引⽤类型，并
发读写时多个协程⻅是通过指针访问同⼀个地址，即访问共享变量，此时同时读写资源
存在竞争关系。会报错误信息:“fatal error: concurrent map read and map write”。
因此，在  Get 中也需要加锁，因为这⾥只是读，建议使⽤读写锁  sync.RWMutex 。
*/
