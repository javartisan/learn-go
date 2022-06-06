package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8099")
	if err != nil {
		fmt.Println("listen on network failed ", err)
		os.Exit(100)
	}

	for true {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept on network failed ", err)
			os.Exit(100)
		}
		go Service(conn)
	}

}

func Service(conn net.Conn) {

	defer conn.Close()

	buf := make([]byte, 1024)
	readLen, _ := conn.Read(buf)
	// 邪恶的转码
	content := string([]rune(string(buf[0:readLen])))
	writeLen, _ := conn.Write(buf[0:readLen])
	fmt.Println("read len:", readLen, "content :", content, "writeLen:", writeLen)
}
