package main

import "fmt"

func main() {

	ch := make(chan int, 2)

	for i := 0; i < 10; i++ {
		// select 当多个case分支同时满足条件时候并不是按照代码顺序执行的，而是随机选择一个case执行。
		//之后结束进行下一次的select执行，也不会case1执行完毕之后 执行下面其他的case
		select {
		case v := <-ch:
			fmt.Println("i ", i, "v = ", v)
			// 第一次写0
		case ch <- i:
			fmt.Println("\t\t  ", "i ", i)
		}
	}

}
