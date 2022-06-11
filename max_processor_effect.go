package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 将CPU参数设置1或者2执行看效果
	runtime.GOMAXPROCS(1)
	for i := 0; i < 100; i++ {
		go fmt.Print(0)
		fmt.Print(1)
	}
	fmt.Println()
}
