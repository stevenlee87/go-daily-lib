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
