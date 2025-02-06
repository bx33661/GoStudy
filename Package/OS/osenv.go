package main

import (
	"fmt"
	"os"
)

func main() {
	//设置环境变量
	os.Setenv("bxname", "bx")

	//获取环境变量
	bxName := os.Getenv("bxname")
	fmt.Println(bxName)
}
