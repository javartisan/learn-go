package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	for i := 0; i < 10; i++ {
		select {
		case v := <-ch:
			fmt.Println("v = ", v)
			// 第一次写0
		case ch <- i:
			fmt.Println("wirte ", i)
		}
	}

}
