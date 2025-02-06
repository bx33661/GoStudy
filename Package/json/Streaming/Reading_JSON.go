package main

import (
	"encoding/json"
	"os"
)

func main() {
	file, _ := os.Open("test.json")
	decoder := json.NewDecoder(file)
	var user []User
}
