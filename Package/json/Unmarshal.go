package main

import (
	"encoding/json"
	"fmt"
)

type User0 struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
	Addr string `json:"address,omitempty"`
}

func main() {
	jsonStr := `{"name":"bx","age":18,"address":"bj"}`
	var user User0
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
