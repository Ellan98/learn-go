package main

import "fmt"

//切片 是可索引的 可以由len()方法获取长度  同时也提供了cap() 可以测量切片最长可以达到多少
// 空切片在未初始前 默认为nil 长度为0
func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
}

func printSlice(numbers []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
}
