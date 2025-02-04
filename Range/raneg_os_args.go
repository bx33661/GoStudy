package main

import (
	"fmt"
	"os"
)

func main() {
	//输出环境变量
	fmt.Println(len(os.Environ()))
	for _, arg := range os.Environ() {
		fmt.Println(arg)
	}
}
