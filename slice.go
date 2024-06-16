package main

import "fmt"

//切片 是对 数组的抽象  由于 数组的长度 不可改变 go中提供了；内置类型的 切片  动态数据
//切片只是对 现有数组 进行引用
//从概念上将 slice 像一个结构体  这个结构体 包含三个元素
//指针   指向数组中slice 指定开始的位置
//长度
//最大长度
// 定义切片 格式如下
//var indentifier []type
//或者使用make函数进行创建
// var slice1 []type = make([]type,len)
//简写 e_slice := make([]type,len)  make([]T,length,capactiy)
func main() {
	e_slice := make([]int, 3)
	e_slice[0] = 1
	e_slice[1] = 2
	e_slice[2] = 3
	e_slice = append(e_slice, 4)
	//前闭后开 s := arr[startIndex: ],startindex 到 endIndex - 1 下的元素 创建为以恶个新的切片 长度为 endIndex - startIndex
	//前开后闭 s:= arr[:endIndex],缺少startIndex  表示冲第一个元素开始
	fmt.Printf("切片%v,其长度为%v", e_slice, len(e_slice))

}
