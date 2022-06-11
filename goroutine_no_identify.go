package main

import (
	"fmt"
	"math/rand"
)

/**

go语言不建议使用thread local
go语言设计者认为代码的行为应该由显示参数决定而不是受一些隐式参数决定
这会导致代码不易读而且会随着goroutine标识导致程序出现故意行为
*/
func main() {

	fmt.Println(rand.ExpFloat64())
}
