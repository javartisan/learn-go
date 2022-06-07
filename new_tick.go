package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ticker.C)
	}

	ticker.Stop()
}
