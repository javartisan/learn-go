package main

import (
	"fmt"
	"sync"
)

// 死锁示范
func main() {

	done := make(chan struct{})
	isCancel := cancel2(done)
	fmt.Println("isCancel = ", isCancel)

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("wg1", wg) // wg1 与 wg2不一样
	go func() {
		fmt.Println("wg2", wg)
		done <- struct{}{} // 此处阻塞
		wg.Done()
	}()

	exit := make(chan struct{})
	fmt.Println("wg1", wg)
	go func() {
		wg.Wait() // 此处阻塞
		isCancel = cancel2(done)
		fmt.Println("isCancel = ", isCancel)
		exit <- struct{}{}
	}()

	<-exit  // 此处阻塞

	// 三处阻塞导致没有活跃的goroutine 最终死锁  那如何解决死锁呢？
	// 解除一个goroutine的阻塞，参考close_go_prog3.go解决
}

// 是否取消
func cancel2(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
