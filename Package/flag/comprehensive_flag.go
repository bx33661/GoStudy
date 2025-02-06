package main

import (
    "flag"
    "fmt"
)

func main() {
    // 1. 定义命令行参数
    // flag.String() 返回一个字符串指针
    name := flag.String("name", "default", "your name")
    // flag.Int() 返回一个整数指针
    age := flag.Int("age", 0, "your age")
    // flag.Bool() 返回一个布尔指针
    verbose := flag.Bool("verbose", false, "enable verbose mode")

    // 2. 解析命令行参数
    // 在使用参数之前必须调用 flag.Parse()
    flag.Parse()

    // 3. 使用参数
    fmt.Printf("Name: %s\n", *name)
    fmt.Printf("Age: %d\n", *age)
    fmt.Printf("Verbose: %v\n", *verbose)

    // 4. flag.Args() 获取非选项参数（即不带 flag 的参数）
    fmt.Println("Non-flag arguments:")
    for i, arg := range flag.Args() {
        fmt.Printf("arg[%d]=%s\n", i, arg)
    }

    // 5. flag.Usage() 打印使用说明
    if *verbose {
        fmt.Println("\nUsage information:")
        flag.Usage()
    }
}