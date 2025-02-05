package main

import "fmt"

func main() {
	Chinese_str := "中华民族万岁"
	// 转换为 rune 切片
	runes := []rune(Chinese_str)

	for i, v := range runes {
		fmt.Printf("位置:%d 字符:%c\n", i, v)
	}
}
