package main

import (
	"fmt"
	"slices"
)

func main() {

	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 3
	slice[2] = 4
	// slices 包中的一个函数，用于检查切片中是否包含特定元素。可以用来替代手动编写循环来检查元素是否在切片中。
	ok := slices.Contains(slice, 3)

	fmt.Printf("当前切片是否存在：%v", ok)

}
