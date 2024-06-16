package main

import "fmt"

// map[key] 获取value 值
//如果key 不存在的时候会  会得到该value值类型的 默认值 但是不会报错
// ok-idiom  获取值 可以知道 key/value 是否存在
func main() {
	// value,ok := map[key]
	m := make(map[string]int)
	m["a"] = 1
	// x, ok := m["b"]
	// fmt.Println(x, ok)
	x, ok := m["a"]

	fmt.Println(x, ok)
}

// 与切片类似 映射是 引用类型当将映射分配 给一个新变量时 它们都指向相同的内部数据结果

//tips
// map 不能使用  = 操作符 进行比较   == 只能用来检查map是否为空 否则会报错
