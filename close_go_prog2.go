//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//
//	done := make(chan struct{})
//	isCancel := cancel(done)
//	fmt.Println("isCancel = ", isCancel)
//
//	var wg sync.WaitGroup
//	wg.Add(1)
//
//	// WaitGroup必须参数传递 否则函数外面引用都是拷贝传递 ，下面是错误示范：
//	//fmt.Println("wg1",wg)  // wg1 与 wg2不一样
//	//go func() {
//	//	fmt.Println("wg2",wg)
//	//	done <- struct{}{}
//	//	wg.Done()
//	//}()
//
//	// 正确示范
//	fmt.Println("wg11", wg)
//	go func(pwg *sync.WaitGroup) {
//		fmt.Println("wg22", pwg)
//		pwg.Done()
//		done <- struct{}{}
//
//	}(&wg)
//
//	wg.Wait()
//	isCancel = cancel(done)
//	fmt.Println("isCancel = ", isCancel)
//}
//
//// 是否取消
//func cancel(ch chan struct{}) bool {
//	select {
//	case <-ch:
//		return true
//	default:
//		return false
//	}
//}
