package main

type Animal1 interface {
	Speak() string
}

type dog0 struct {
}

type cat0 struct {
	name string
}

func (PiPi dog0) Speak() string {
	return "汪汪汪"
}

func (XiaoHua cat0) Speak() string {
	return "喵喵喵"
}

func main() {
	demo := cat0{name: "猫"}
	println(demo.Speak())
}
