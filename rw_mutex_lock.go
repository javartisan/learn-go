package main

import (
	"fmt"
	"sync"
)

var rwLock sync.RWMutex

var ch = make(chan int, 20)

func main() {
	for i := 0; i < 10; i++ {
		go write()
		go read()
	}

	for i := 0; i < 20; i++ {
		<-ch
	}

	rwLock.Lock()
	//rwLock.Lock()
	//  RWMutex是不可以重入锁
	//rwLock.Unlock()
	rwLock.Unlock()

}

func read() {
	rwLock.RLock()
	fmt.Println("read start")
	defer rwLock.RUnlock()
	fmt.Println("read done")
	ch <- 1
}

func write() {
	rwLock.Lock()
	fmt.Println("write start")
	defer rwLock.Unlock()
	fmt.Println("write done")
	ch <- 1
}
