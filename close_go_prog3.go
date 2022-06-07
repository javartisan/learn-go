package main

import (
	"fmt"
	"sync"
)

// 死锁示范
func main() {

	done := make(chan struct{}, 1) // 一个缓冲的chan
	isCancel := cancel3(done)
	fmt.Println("isCancel = ", isCancel)

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("wg1", wg) // wg1 与 wg2不一样
	go func() {
		fmt.Println("wg2", wg)
		done <- struct{}{} // 注意：此处将不再阻塞
		wg.Done()
	}()

	exit := make(chan struct{})
	fmt.Println("wg1", wg)
	go func() {
		wg.Wait() // 此处阻塞
		isCancel = cancel3(done)
		fmt.Println("isCancel = ", isCancel)
		exit <- struct{}{}
	}()

	<-exit // 此处阻塞

	//  将21行代码放行最终让所有goroutine执行完毕
}

// 是否取消
func cancel3(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
