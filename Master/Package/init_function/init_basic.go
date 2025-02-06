package main

import "fmt"

var count int

// init()函数会在main()函数之前执行，且只会执行一次。
// init()函数的作用是包的初始化，比如初始化包的变量等。
func init() {
	count = 42
	fmt.Println("init() function called")
}

func main() {
	fmt.Println("main() function called")
	fmt.Println("count:", count)
}
