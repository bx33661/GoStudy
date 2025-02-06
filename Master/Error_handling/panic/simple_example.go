package main

import "fmt"

// panic函数的作用是引发一个运行时错误，这个错误会引起程序的崩溃
func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred!!!!")
	fmt.Println("Ending the program")
	fmt.Println("这个不会执行的")
}
