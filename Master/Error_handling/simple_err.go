package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建一个错误,调用errors包的New函数
	err := errors.New("This is an error!!!!!!")
	fmt.Print(err)
}
