package main

import "fmt"

// 定义接口
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪汪"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵喵"
}

// 多态的使用
func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	animals := []Speaker{
		Dog{},
		Cat{},
	}

	for _, animal := range animals {
		MakeSpeak(animal)
	}
}
