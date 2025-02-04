package main

import "fmt"

func main() {
	Chinese_str := "中华民族万岁"
	fmt.Print(Chinese_str)
	for i, v := range Chinese_str {
		fmt.Print(i, Chinese_str[v-1], " ")
	}
}
