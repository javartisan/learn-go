package main

import (
	"fmt"
	"time"
)

// 无缓冲的chan  提供强同步元语

func main() {
	ch := make(chan int)
	// 存在goroutine泄露的问题  因为调用了三次 ch写  由于两个goroutine无法写成功 因此内存泄露
	for i := 0; i < 3; i++ {
		go func() {
			ch <- 1
		}()
	}

	fmt.Println(<-ch)

	time.Sleep(10 * time.Second)
}
