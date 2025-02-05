package main

import "fmt"

func main() {
	//Go语言中的数据类型转换只能在两种相互兼容的类型之间进行
	//GO不支持隐式类型转换
	var a int64 = 3
	var b int32
	b = a
	fmt.Printf("b 为 : %d", b)
}
