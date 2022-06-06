package main

import "fmt"

func main() {

	ints := make(chan int, 10)

	for i := 0; i < 15; i++ {
		ints <- i
		fmt.Println("cap chan = ", cap(ints), "len chan = ", len(ints))
	}
}
