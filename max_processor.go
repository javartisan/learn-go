package main

import (
	"fmt"
	"runtime"
)

/**
Goroutines that are sleeping or blocked in a communication do not need a thread at all.
Goroutines that are blocked in I/O or other system calls or are calling non-Go functions,
do need an OS thread, but GOMAXPROCS need not account for them.


Goroutine的休眠以及通道通信阻塞是不需要占用线程的。

但是如果IO阻塞或系统调用或者调用非Go函数则会占用一个操作系统线程，但是被占用的系统线程不会包括在GOMAXPROCS里面
*/
func main() {

	cpu := runtime.NumCPU()
	fmt.Println("NumCPU = ", cpu)

}


