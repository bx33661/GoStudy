# flag库



---

**flag 包**的主要特点：

1. 支持的数据类型：
   - String
   - Int, Int64
   - Uint, Uint64
   - Float64
   - Bool
   - Duration
2. 参数格式：
   - `-flag`
   - `-flag=x`
   - `-flag x`
3. 常用函数：
   - `flag.String()`, `flag.Int()` 等：定义新的命令行参数
   - `flag.Parse()`：解析命令行参数
   - `flag.Usage()`：打印使用说明
   - `flag.Args()`：获取非选项参数
4. 自动生成帮助：
   - 使用 `-h` 或 `-help` 可以自动显示帮助信息



```go
s := flag.String("s", "default", "字符串说明")
```

- `"-s"`：定义了一个**短选项**（`-s`），用户在运行程序时可以通过 `-s <value>` 传入值。
- `"default"`：设置该参数的**默认值**，如果用户没有传入 `-s` 选项，则 `s` 的值为 `"default"`。
- `"字符串说明"`：这是该选项的**帮助信息**，当用户运行 `your_program -h` 时，它会显示这条说明。



这里给出一个`example.go`

```go
package main

import (
	"flag"
	"fmt"
)

func main(){
	name := flag.String("name","default","输入名称")
	age := flag.Int("age",0,"输入年龄")
	isVerbose := flag.Bool("verbose",false,"是否显示详细信息")

	flag.Parse()

	fmt.Printf("名称：%s\n",*name)
	fmt.Printf("年龄：%d\n",*age)
	fmt.Printf("详细信息: %v\n",*isVerbose)

}
```

> 我这里是Windows电脑：exe
> ```go
> go build -o flagdemo.exe example.go
> ./flagdemo.exe -name="Alice" -age=25 -verbose=True
> ```

### 参数格式

就是说三种格式都可以使用

1. 使用 `-flag=x` 格式：

```bash
./flagdemo.exe -name="Alice" -age=25 -verbose=True
```

2. 使用 `-flag x` 格式：

```bash
./flagdemo.exe -name "Alice" -age 25 -verbose True
```

3. 使用 `-flag` 格式

```bash
# 仅使用 -flag 就表示将该标志设置为 true
go run flag_demo.go -verbose -debug
```



## 常用函数

### `flag.Usage()`函数

> `flag.Usage()` 会在以下情况下自动调用：
>
> - 当用户使用 `-h` 或 `-help` 标志时
> - 当命令行参数解析出错时

**自定义 Usage 信息**（可选）：

- 可以通过重写 `flag.Usage` 函数来自定义帮助信息的格式
- `flag.PrintDefaults()` 会打印所有已定义参数的说明

示例函数

```go
package main

import (
    "flag"
    "fmt"
)

func main() {
    // 1. 定义命令行参数
    name := flag.String("name", "world", "specify a name to greet")
    age := flag.Int("age", 0, "specify person's age")
    verbose := flag.Bool("verbose", false, "enable verbose output")

    // 2. 自定义 Usage 信息（可选）
    flag.Usage = func() {
        fmt.Printf("Usage of %s:\n", "your_program_name")
        fmt.Printf("Example: your_program_name -name Alice -age 25\n\n")
        fmt.Println("Available flags:")
        flag.PrintDefaults()
    }

    // 3. 解析命令行参数
    flag.Parse()

    // 4. 使用参数
    if *verbose {
        fmt.Printf("Name: %s, Age: %d\n", *name, *age)
    } else {
        fmt.Printf("Hello, %s!\n", *name)
    }
}
```

|||||||

**程序输出**如下----

```go
PS G:\go\awesomeProject\Package\flag> ./flagusage.exe -h
Usage of your_program_name:
Example: your_program_name -name Alice -age 25

Available flags:
  -age int
        specify person's age
  -name string
        specify a name to greet (default "world")
  -verbose
        enable verbose output
```

