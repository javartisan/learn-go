package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})

	go func() {
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			fmt.Println("err : ", err)
			return
		}
		fmt.Println("abort <- struct{}{}")
		abort <- struct{}{}
	}()

	// 类似于一个定时的chan
	tickChan := time.Tick(1 * time.Second)
	for count := 0; count < 5; count++ {
		fmt.Println(count, " <-tickChan : ", <-tickChan)
	}

	now := time.Now()
	after := time.After(2 * time.Second)

	//select {
	//// 空select将会导致永远等待
	//}

	select {
	case val := <-after:
		fmt.Println("do nothing at after")
		fmt.Println("now : ", now, "chan after : ", val)
	case <-abort:
		fmt.Println("abort launch")
		os.Exit(100)
	}

	launch()

}

func launch() {
	fmt.Println("launch...")
}
