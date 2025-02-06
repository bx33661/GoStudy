# Go中的错误处理

---

Go语言采用了一种独特的错误处理方式，它不使用传统的try-catch异常处理，而是通过返回值来进行错误处理。

一个典型的例子：

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
       return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(3, 0)
    if err != nil {
       fmt.Println(err)
    } else {
       fmt.Println(result)
    }
}
```

这里有一个`nil`值，之前我们在切片等地方也看见过

**`nil` 是 Go 语言中的一个特殊值，代表某些类型的"零值"或"空值"。它类似于其他语言中的 `null`，但有其独特的特点。**



## Go报错分析

```bash
G:\go\awesomeProject\Master\Error_handling\panic git:[main]
go run simple_example.go
Starting the program
panic: A severe error occurred!!!!

goroutine 1 [running]:
main.main()
        G:/go/awesomeProject/Master/Error_handling/panic/simple_example.go:8 +0x59
```

程序内容输出,错误消息

```bash
Starting the program
panic: A severe error occurred!!!!
```

**goroutine 1** 表示错误发生在主协程（程序的主线程）

**main.main()** 指出错误发生在 `main` 包中的 `main` 函数

**simple_example.go:8** 明确指出了错误发生的代码位置：在`simple_example.go`文件第 8 行

**+0x59** 是内存偏移地址，*通常不需要关注*

exit status 2，表示程序因错误非正常退出，状态码为 2。



## Panic,Recover,defer

### 执行顺序

```go
package main

import "fmt"

func main() {
	fmt.Println("1. 进入 main 函数")

	// 注册第一个 defer
	defer fmt.Println("6. 第一个 defer")

	fmt.Println("2. main 函数继续执行")

	// 注册第二个 defer
	defer fmt.Println("5. 第二个 defer")

	fmt.Println("3. 即将发生 panic")

	// 注册第三个 defer（包含 recover）
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("4. recover 捕获到 panic:", r)
		}
	}()

	panic("发生 panic！")

	// 下面的代码永远不会执行
	fmt.Println("这行不会执行")
}
```

1. 正常代码就顺序执行
2. 但是`panic`触发之后，按照 LIFO（后进先出）的顺序执行 defer 函数
3. **recover 处理**，如果 defer 函数中有 recover()，可以捕获 panic，recover 之后，程序可以继续执行剩余的 defer 函数，如果没有 recover，程序会在执行完所有 defer 后终止

> 这种机制确保了在程序发生严重错误时，仍然能够执行必要的清理工作（通过 defer），同时提供了恢复机制（通过 recover）。
