package main

import "fmt"

func main() {
	defer fmt.Println("1. deferred call in main")
	defer fmt.Println("2. defer 执行")

	fmt.Println("Main function")
	panic("A severe error occurred")
	fmt.Println("Main function")
	// 输出:
	// 正常语句
	// 2. defer 执行
	// 1. defer 执行
	// panic: 发生 panic！
}
