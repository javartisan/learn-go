package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var count = 0

func main() {

	for true {
		Request()
		time.Sleep(1 * time.Second)
	}

}

func Request() {
	conn, err := net.Dial("tcp", "localhost:8099")
	if err != nil {
		fmt.Println("dial on network failed ", err)
		os.Exit(100)
	}
	defer conn.Close()

	writeLen, _ := conn.Write([]byte("你好，Golang:" + strconv.Itoa(count)))
	buf := make([]byte, 1024)

	read, _ := conn.Read(buf)

	content := string([]rune(string(buf[0:read])))

	fmt.Println("count = ", count, "write len:", writeLen, "content:", content, "read len:", read)
	count++

}
