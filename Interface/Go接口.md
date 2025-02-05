# Go接口学习与分析

---

> Go通过这些方式提供了面向对象语言的类似功能

写了几个例子：

```go
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

```

细化分析学习一下，先看这个

```go
func (b Bird) speak() string {
    return "叽叽叽"
}
```

 Go 语言中的方法声明语法,`(b Bird)`是方法接收器**（Method Receiver）**，

- `b` 是接收器变量名（可以在方法内使用这个变量来访问结构体的属性）
- `Bird` 是接收器类型（表示这个方法属于 Bird 结构体）

`speak()`就是方法名了，`string`即是方法的返回值类型



## 实现面向对象的方式

### 封装



### 继承的实现

```go
package main
// 采用内嵌的方式实现继承，组合的方式
import "fmt"

type Animal struct {
	Name string
	age  int
}

// 采用内嵌的方式实现继承
type Dog struct {
	Animal
	breed string
}

func main() {
	// 初始化Dog对象,并赋值,这里的Name和age是Animal的属性
	dog := Dog{
		Animal: Animal{
			Name: "刘富贵",
			age:  2,
		},
		breed: "哈士奇",
	}

	//可以直接访问Animal的属性
	fmt.Println(dog.Name)
	fmt.Println(dog.age)
}

```

