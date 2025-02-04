package main

import "fmt"

func main() {
	//创建一个多维数组
	var nums [5][5]int
	//为每个数组赋值
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			nums[i][j] = nums[i][j] + 1
		}
	}

	//打印整个数组
	fmt.Println(nums)

	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println(a)
}
