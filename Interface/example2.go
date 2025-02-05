package main

import "fmt"

type Animal interface {
	speak() string
}

type dog struct{}

func (d dog) speak() string {
	return "汪汪汪"
}

type cat struct{}

func (c cat) speak() string {
	return "喵喵喵"
}

type Bird struct{}

func (b Bird) speak() string {
	return "叽叽叽"
}

func AllSpeak(a []Animal) {
	for _, v := range a {
		fmt.Print(v.speak())
	}
}

func main() {
	aimls := []Animal{dog{}, cat{}, Bird{}}
	AllSpeak(aimls)
}
