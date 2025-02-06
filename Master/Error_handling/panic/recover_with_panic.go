package main

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Printf("恢复自 panic: %v\n", r)
	}
}

func mayPanic() {
	defer handlePanic()

	// 引发 panic
	panic("发生了错误❌")
}

func main() {
	fmt.Println("开始")
	mayPanic()
	fmt.Println("结束") // 这行会执行，因为 panic 被恢复了
}
