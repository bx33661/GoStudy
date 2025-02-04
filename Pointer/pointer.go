package main

import "fmt"

//指针的零值是 nil

func main() {
	//一个指针的简单演示
	a := 10
	var p *int
	p = &a
	fmt.Println(p)
	fmt.Println(*p)

	//通过指针修改值
	x := 20
	y := &x
	fmt.Println(*y)
	modify(y)
	fmt.Println(*y)
}

func modify(ptr *int) {
	*ptr = 100
}
