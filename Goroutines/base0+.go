package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i:", i)
		}(i)
		//这里加上time.Sleep(time.Millisecond)是为了让goroutine有机会执行，否则主goroutine执行完毕后，程序就退出了。
		time.Sleep(time.Millisecond)
	}
}
