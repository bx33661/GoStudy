package main

// 模仿一个例子
type Base interface {
	//定义一个方法
	Show()
}

type BaseImpl struct {
	radius int
}

func (c BaseImpl) Show() {
	println("BaseImpl Show")
}

func main() {
	var b Base
	b = BaseImpl{
		radius: 10,
	}
	b.Show()
}
