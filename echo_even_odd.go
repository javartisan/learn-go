package main

import (
	"fmt"
	"time"
)

func main() {
	// 无缓冲通道
	even := make(chan int)
	// 等效如上写法，无缓冲通道
	odd := make(chan int, 0)

	data := make(map[string]string)
	data["m1"] = "m1"
	fmt.Println(data)
	Put(data)
	fmt.Println(data)

	slice := make([]int, 0, 10)
	slice = append(slice, 1)
	fmt.Println(slice)
	slice = Add(slice)
	fmt.Println(slice)

	go Odd(even, odd)
	go Even(even, odd)
	// 启动
	// odd <- 1

	time.Sleep(10 * time.Second)
}

// 分片传递的是值
func Add(slice []int) (s []int) {
	s = append(slice, 2)
	return s
}

// map传递的是引用
func Put(m map[string]string) {
	m["m"] = "m"
}

func Odd(even, odd chan int) {
	for true {
		val := <-odd
		fmt.Println("Odd:", val)
		even <- val + 1
	}
}

func Even(even, odd chan int) {
	for true {
		val := <-even
		fmt.Println("Even:", val)
		odd <- val + 1
	}
}
