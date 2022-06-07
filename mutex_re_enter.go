package main

import "sync"

func main() {

	//  sync.Mutex 是不可重入的锁
	var lock sync.Mutex

	lock.Lock()
	lock.Lock()
	defer lock.Unlock()
	defer lock.Unlock()
}
