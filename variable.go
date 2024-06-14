package main

import "fmt"

// var name type    声明后若不赋值，使用默认值
// var name type = value   格式
// var name  = value  根据值自行判定变量类型(类型推断Type inference)
// name := value 左侧的变量不应该是已经声明过的  只能被用在函数体内，而不可以用于全局变量的声明与赋值
// 多变量声明
// var name1, name2, name3 type 以逗号分隔，声明与赋值分开，若不赋值，存在默认值
// var name1, name2, name3 string = 1, v2, v3
// var name1, name2, name3 = v1, v2, v3  直接赋值，下面的变量类型可以是不同的类型
// var (
// 	name1 type1
// 	name2 type2
// )
var name string = "value"

func main() {
	fmt.Println(name)
}

//tips
// 变量必须先定义才能使用
// go语言是静态语言，要求变量的类型和赋值的类型必须一致。
// 变量的零值。也叫默认值。
// 变量定义了就要使用，否则无法通过编译。
