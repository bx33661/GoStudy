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