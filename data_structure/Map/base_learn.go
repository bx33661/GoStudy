package main

import "fmt"

func main() {
	//第一种
	ages0 := make(map[string]int)
	ages0["a"] = 10
	ages0["b"] = 20
	ages0["c"] = 30

	//第二种
	ages1 := map[string]int{
		"James": 40,
		"AD":    32,
	}

	for name, age := range ages1 {
		fmt.Println(name, age)
	}

	for key := range ages1 {
		fmt.Println(key)
	}

	for _, age := range ages0 {
		fmt.Println(age)
	}

	//获取Map的长度
	len := len(ages0)
	fmt.Print(len)

	//删除元素
	delete(ages0, "a")
}
