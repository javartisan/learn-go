package main

import (
	"fmt"
	"gopl.io/ch8/thumbnail"
)

func main() {

	filePath := "/Users/daxin/Desktop/t1.jpg"
	file, err := thumbnail.ImageFile(filePath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(file)

}
