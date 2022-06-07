package main

import (
	"fmt"
	"sync"
	"time"
)

// Once是一个互斥锁与一个bool值封装而成的
var loadOnce sync.Once

func main() {

	for i := 0; i < 10; i++ {
		// 仅仅执行一次
		go loadOnce.Do(loadData)
	}
	time.Sleep(time.Second)
}

func loadData() {
	fmt.Println("loadData done")
}
