package main

import (
	"fmt"
)

// Go中 的字符串 是一个字节的切片Unicode 兼容的  并且是UTF-8编码的
// 访问strings 包可以操作很多string的函数
func main() {
	var test string = "hello world"
	fmt.Println(test)
	for i := 0; i < len(test); i++ {
		fmt.Printf("当前索引为%d   Unicode为%v , 值为%c \n", i, test[i])
	}
}
