package main

import "fmt"

func main() {
	Chinese_str := "中华民族万岁"
	fmt.Println(Chinese_str)

	// 使用 range 遍历字符串
	for i, v := range Chinese_str {
		fmt.Printf("索引:%d 字符:%c Unicode:%d\n", i, v, v)
	}
}
