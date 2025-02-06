package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("hello python")
	err := os.WriteFile("test.txt", data, 0644)
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
