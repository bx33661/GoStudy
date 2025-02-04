package main

import "fmt"

func main() {
	x, y := 3, 4
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

/*
var p *int
*p = 100  // 错误！会导致panic，因为p是nil

// 正确的做法
var x int
p = &x
*p = 100
*/
