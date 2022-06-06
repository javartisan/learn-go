package main

import (
	"fmt"
	"time"
)

func main() {

	natural := make(chan int32)

	go func() {
		n := <-natural
		fmt.Println("read ch : ", n)
	}()
	time.Sleep(2 * time.Second)
	close(natural)
	fmt.Println("close(natural) done")
	time.Sleep(2 * time.Second)
}
