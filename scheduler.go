package main

import (
	"fmt"
	"time"
)

/**
1: 操作系统线程由硬件时钟终端发送到CPU然后调用内核的调度函数选择一个线程进行执行，
	这个过程会保存当前线程的状态信息，然后将被执行的线程信息进行恢复（上下文切换）

2：Go运行时拥有自己的调度器，这个调度器使用的是m:n调度的技术，即将m个goroutine调度到n个OS线程上执行。
	Go调度器与内核调度器工作类似，但是Go调度器只负责调度Go程序内部的Goroutine

3：Go调度器实现不是由硬件时钟中断，而是软件实现调度。当一个Goroutine调用time.Sleep或被通道阻塞或对互斥量操作时，
	Go运行时调度器会将这个Goroutine设置为休眠模式，并运行其他的Goroutine直到前一个可以被重新唤醒为止

4：由于Goroutine的调度是在用户态完成的，不需要操作系统内核干预  所以效率较高


重点：
Goroutines that are sleeping or blocked in a communication do not need a thread at all.
Goroutines that are blocked in I/O or other system calls or are calling non-Go functions,
do need an OS thread, but GOMAXPROCS need not account for them.

*/

func main() {

	start := time.Now().UnixNano()
	end := start + time.Second.Nanoseconds()

	count := 0

	pc := &count

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for time.Now().UnixNano() <= end {
			val := <-ch1
			ch2 <- val
			*pc++
		}
	}()

	go func() {
		for time.Now().UnixNano() <= end {
			val := <-ch2
			ch1 <- val
		}
	}()

	ch1 <- 1

	for time.Now().UnixNano() <= end {

	}

	fmt.Println("costime = ", end-start, "nano count =", *pc)

}
