package main

import "fmt"

//append向slice内最佳一个或者多个元素  返回 一个和额slice一样类型的slice
//append 会改变slice所引用的数组内容 从而影响到引用同一数组的其它slice
//当slice中没有剩余空间 即cap-len == 0，此时 将动态分配新的数组空间，返回的slice数组指针将指向这个空间而原数组内容将保持不变 其它引用此数组的slice则不受影响

//  copy函数 从源slice的src中复制元素到目标dst并且返回复制的元素个数

func main() {
	var numbers []int
	printScliec(numbers)
	// 允许追加空切片
	numbers = append(numbers, 0)
	printScliec(numbers)
	// 向切片添加一个元素
	numbers = append(numbers, 1)
	printScliec(numbers)
	// 同时添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printScliec(numbers)
	// 创建切片 numbers1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	//拷贝numbers 的内容 到numbers1
	copy(numbers1, numbers)
	printScliec(numbers1)

}

func printScliec(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}
