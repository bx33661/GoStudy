package main

import "fmt"

func main() {
	var i int = 36
	var f float64 = 36.0
	var st string = "hello bx"
	var b bool = true
	var c complex64 = 5 + 5i

	//打印类型,使用%T，打印变量的类型
	fmt.Printf("这个类型是：%T\n", i)
	fmt.Printf("这个类型是：%T\n", f)
	fmt.Printf("这个类型是：%T\n", st)
	fmt.Printf("这个类型是：%T\n", b)
	fmt.Printf("这个类型是：%T\n", c)
}
