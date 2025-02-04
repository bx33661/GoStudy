package main

import "fmt"

type Animal struct {
	Name           string
	favoriteFruits string
	id             int
}

func main() {
	//赋值
	var wolf Animal
	wolf.Name = "Wolf"
	wolf.favoriteFruits = "Apple"
	wolf.id = 1
	fmt.Println(wolf)

	//结构体作为函数参数
	echo_animal(wolf)
}

func echo_animal(animal Animal) {
	fmt.Println("最喜欢的食物是： ")
	fmt.Println(animal.favoriteFruits)
}
