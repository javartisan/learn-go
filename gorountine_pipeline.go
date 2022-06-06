package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var closed = false

func main() {

	natural := make(chan int32)
	square := make(chan int32)
	count := 1

	go func() {
		for true {
			mutex.Lock()
			if !closed {
				natural <- int32(count)
				count++
				mutex.Unlock()
				continue
			}
			mutex.Unlock()
			break
		}
	}()

	go func() {
		for true {
			n, ok := <-natural
			if ok {
				square <- n * n
			} else {
				close(square)
				break
			}
		}

	}()

	func() {
		for true {
			n, ok := <-square
			if ok {
				fmt.Println("square = ", n)
			}
			if n > 100 {
				mutex.Lock()
				if !closed {
					closed = true
					close(natural)
					break
				}
				mutex.Unlock()
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
