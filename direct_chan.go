package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go SendChan(ch)
	go RevChan(ch)
	time.Sleep(1 * time.Second)
}

func SendChan(send chan<- int) {
	send <- 1
	fmt.Println("SendChan done ")
}
func RevChan(rev <-chan int) {
	val := <-rev
	fmt.Println("RevChan = ", val)
}
