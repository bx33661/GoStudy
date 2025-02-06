package main

import (
	"fmt"
	"os"
)

func main() {
	//获取文件 目录的元信息（如大小、修改时间）
	fileInfo, err := os.Stat("test.txt")
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件大小:", fileInfo.Size(), "修改时间:", fileInfo.ModTime())
		fmt.Println("文件是否为目录:", fileInfo.IsDir())
	}

	//获取当前目录
	dir, _ := os.Getwd()
	fmt.Println("当前目录:", dir)

}
