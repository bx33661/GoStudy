package main

import (
	"encoding/json"
	"fmt"
)

//json序列化

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}

func main() {
	user := User{
		Name: "bx",
		Age:  18,
		Addr: "bj",
	}
	//序列化
	date, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(date)
	fmt.Println(string(date))
	/* result:
	[123 34 110 97 109 101 34 58 34 98 120 34 44 34 97 103 101 34 58 49 56 44 34 97 100 100 114 34 58 34 98 106 34 125]
	{"name":"bx","age":18,"addr":"bj"}
	*/
}
