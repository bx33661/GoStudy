package main

import (
	"flag"
	"fmt"
)

func main(){
	name := flag.String("name","default","输入名称")
	age := flag.Int("age",0,"输入年龄")
	isVerbose := flag.Bool("verbose",false,"是否显示详细信息")

	//必须在使用参数之前调用
	flag.Parse()

	
	fmt.Printf("名称：%s\n",*name)
	fmt.Printf("年龄：%d\n",*age)
	fmt.Printf("详细信息: %v\n",*isVerbose)

}
