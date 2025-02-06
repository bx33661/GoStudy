package main

import (
	"fmt"
	"sync"
)

func main() {
	//sync.WaitGroup是一个计数信号量，用来记录goroutine的完成。
	var wg sync.WaitGroup
	//但是这个程序并不会按顺序输出0-9，因为goroutine是并发执行的，所以输出的顺序是随机的。
	for i := 0; i < 10; i++ {
		//Add(1)表示计数器+1，Done()表示计数器-1，Wait()表示等待计数器为0。
		wg.Add(1)
		go func(i int) {
			//defer表示延迟执行，这里表示在goroutine结束时执行wg.Done()。
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	//等待所有goroutine执行完毕
	wg.Wait()
}
