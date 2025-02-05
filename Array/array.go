package main

import "fmt"

func main() {
	//初始化数组
	var n [10]int
	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}
	for i := 0; i < 10; i++ {
		fmt.Println("值为:", n[i])
	}

	//声明数组并切赋值
	bili := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(bili)
}
