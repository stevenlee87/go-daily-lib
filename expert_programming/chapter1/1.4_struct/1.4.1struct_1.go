package main

type People struct {
	Name string
	age  int
}

func (p *People) SetName(name string) {
	p.Name = name
}

func (p *People) SetAge(age int) {
	p.age = age
}

func (p *People) GetAge() int {
	return p.age
}

func main() {
	p := People{Name: "lee", age: 18}
	p.GetAge()
}
