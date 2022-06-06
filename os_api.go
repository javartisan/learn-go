package main

import (
	"fmt"
	"os"
)

func main() {

	for index, value := range os.Environ() {
		fmt.Println(index, value)
	}

	env, ok := os.LookupEnv("MVN_HOME")
	if ok {
		fmt.Println("\t", "\t", "\t", "\t", "env = ", env)
	}
}
