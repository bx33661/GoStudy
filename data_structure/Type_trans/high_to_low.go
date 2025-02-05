package main

import (
	"fmt"
)

func main() {
	var i64 int64 = 123456789012345
	var f64 float64 = 123456789.123456789

	var i int = int(i64) // 从 int64 转换为 int，可能会丢失数据
	var f float32 = float32(f64)
	// 从 float64 转换为 float32，可能会丢失精度

	fmt.Printf("i: %d, f: %f\n", i, f)
}
