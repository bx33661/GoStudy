package main

// 采用内嵌的方式实现继承，组合的方式
import "fmt"

type Animal struct {
	Name string
	age  int
}

// 采用内嵌的方式实现继承
type Dog0 struct {
	Animal
	breed string
}

func main() {
	// 初始化Dog对象,并赋值,这里的Name和age是Animal的属性
	dog := Dog0{
		Animal: Animal{
			Name: "刘富贵",
			age:  2,
		},
		breed: "哈士奇",
	}

	//可以直接访问Animal的属性
	fmt.Println(dog.Name)
	fmt.Println(dog.age)
}
