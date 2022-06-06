package main

import "fmt"

func main() {

	list := make([]int, 0, 10)
	list = append(list, 1, 2, 3, 4, 5)
	fmt.Println(list)
	list = remove(list, 2)
	fmt.Println(list)
	list = add(list, 3, 2)
	fmt.Println(list)
	list = add(list, len(list)+1, len(list))
	fmt.Println(list)
	list = add(list, 0, 0)
	fmt.Println(list)
}

func remove(list []int, index int) []int {
	if len(list) <= index {
		return list
	}
	return append(list[0:index], list[index+1:]...)
}

//  这种方式有内存创建与拷贝 性能不是最佳
// 可以先尾部添加一个0值之后进行拷贝
func add(list []int, e, index int) []int {
	if len(list) < index {
		return list
	}

	before := list[0:index]
	after := list[index:]
	newList := make([]int, 0, len(list)+1)
	newList = append(newList, before...)
	newList = append(newList, e)
	newList = append(newList, after...)
	return newList
}
