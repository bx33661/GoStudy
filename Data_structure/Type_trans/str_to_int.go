package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "123"

	// Atoi()函数可以将字符串转换为整数,ASCII码转换为Int
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("i type is %T, i value is %d\n", i, i)
	}

	// Itoa()函数可以将整数转换为字符串,Int转换为ASCII码
	s = strconv.Itoa(i)
	fmt.Println("s type is %T, s value is %s\n", s, s)
}
