package main

import (
	"fmt"
	"sync"
)

func main() {

	done := make(chan struct{})
	isCancel := cancel(done)
	fmt.Println("isCancel = ", isCancel)

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("wg11", wg)
	go func(pwg *sync.WaitGroup) {
		fmt.Println("wg22", pwg)
		pwg.Done()
		done <- struct{}{}

	}(&wg)

	exit := make(chan struct{})
	go func() {
		wg.Wait()
		isCancel = cancel(done)
		fmt.Println("isCancel = ", isCancel)
		exit <- struct{}{}
	}()

	<-exit

}

// 是否取消
func cancel(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
