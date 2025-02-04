package main

import "fmt"

func main() {
	goMap := map[string]int{}
	goMap["a"] = 1
	goMap["b"] = 2
	value := goMap["a"]

	//检查键是否存在
	value, exists := goMap["c"]
	if exists {
		fmt.Printf("The value of c is %d\n", value)
	} else {
		fmt.Printf("The value of c is not %d\n", value)
	}

}
