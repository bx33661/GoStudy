package main

//封装的实现
//
type Person struct {
	//首字母小写，外部包不可访问，首字母大写，外部包可访问
	Name string
	age  int
}

//实现Getter的效果
func (p Person) GetAge() int {
	return p.age
}

//实现Setter的效果
func (p *Person) SetAge(age int) {
	p.age = age
}

func main() {

}
