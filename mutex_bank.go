package main

import (
	"fmt"
	"sync"
)

var (
	lock    sync.Mutex
	balance int
)
var ints = make(chan int, 100)

func Deposit(amount int) {
	lock.Lock()
	defer lock.Unlock()
	balance += amount
	ints <- amount

}

func Balance() int {
	lock.Lock() // 获取lock之后立马注册defer lock.Unlock()是一个好习惯
	defer lock.Unlock()
	return balance
}

func main() {

	for i := 0; i < 100; i++ {
		go Deposit(1)

	}

	// 等待计算完成
	for i := 0; i < 100; i++ {
		<-ints
	}

	fmt.Println(Balance())
}
