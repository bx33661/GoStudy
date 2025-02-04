package main

import (
	"fmt"
	"unicode"
)

func main() {
	//Uniocde库使用
	var str rune = '中'
	fmt.Println("是否为汉字？", unicode.Is(unicode.Han, str))

	fmt.Println(string((unicode.ToLower('A'))))
}
