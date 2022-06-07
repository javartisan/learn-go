package main

import (
	"fmt"
	"sync"
)

// 死锁示范
func main() {

	done := make(chan struct{}, 1) // 一个缓冲的chan
	isCancel := cancel4(done)
	fmt.Println("isCancel = ", isCancel)

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("wg1", wg) // wg1 与 wg2不一样
	go toDone(wg, done)

	exit := make(chan struct{})
	fmt.Println("wg1", wg)
	go func() {
		wg.Wait() // 此处阻塞
		isCancel = cancel4(done)
		fmt.Println("isCancel = ", isCancel)
		exit <- struct{}{}
	}()

	<-exit // 此处阻塞

	//  将21行代码放行最终让所有goroutine执行完毕
}

// WaitGroup值拷贝，会跟调用者处的WaitGroup隔离，导致死锁  因此WaitGroup需要使用指针传递 参见close_go_prog5.go
func toDone(wg sync.WaitGroup, done chan struct{}) {
	fmt.Println("toDone", wg)
	done <- struct{}{} // 注意：此处将不再阻塞
	wg.Done()
}

// 是否取消
func cancel4(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
