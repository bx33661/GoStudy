package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	// 长度
	fmt.Println(len(nums))

	// 切片容量
	fmt.Println(cap(nums))

	// 切片操作
	fmt.Println(nums[:2])
	fmt.Println(nums[2:4])

	/*
		go run array_slice.go
		[1 2 3 4 5]
		5
		5
		[1 2]
		[3 4]
	*/
}
