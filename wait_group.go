package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			wg.Done()
		}()
	}

	fmt.Println("main wait")
	wg.Wait()
	fmt.Println("main done")

}
