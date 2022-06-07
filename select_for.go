package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case v := <-ch:
			fmt.Println("v = ", v)
		case ch <- i:
		}
	}

}
