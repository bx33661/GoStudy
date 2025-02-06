package main

import "fmt"

func main() {
	fmt.Println("1. 进入 main 函数")

	// 注册第一个 defer
	defer fmt.Println("6. 第一个 defer")

	fmt.Println("2. main 函数继续执行")

	// 注册第二个 defer
	defer fmt.Println("5. 第二个 defer")

	fmt.Println("3. 即将发生 panic")

	// 注册第三个 defer（包含 recover）
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("4. recover 捕获到 panic:", r)
		}
	}()

	panic("发生 panic！")

	// 下面的代码永远不会执行
	fmt.Println("这行不会执行")
}
