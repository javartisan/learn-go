package main

import "fmt"

func main() {
	call()
	fmt.Println("main done")
}

func call() {

	defer func() {
		if err := recover(); err != nil { // recover检索栈的值
			fmt.Println(err)
		}
	}()

	div(0, 0)
}

func div(a, b int) {
	val := a / b
	fmt.Println(val)
}
