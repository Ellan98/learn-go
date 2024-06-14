package main

import "fmt"

// 数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映

// 数组类型的数据结构。 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，
// 这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
// 可以通过索引（位置）来读取（或者修改），索引从0开始 到长度减1

//格式 声明和初始化数组
// var variable_name [SIZE] variable_type
// var balance [10]float32
// var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// 1.初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
// 2.忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
// var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// balance[4] = 50.0
func main() {
	var a [4]float32 // 等价于：var arr2 = [4]float32{}
	fmt.Println(a)   // [0 0 0 0]
	var b = [5]string{"ruby", "王二狗", "rose"}
	fmt.Println(b)                          // [ruby 王二狗 rose  ]
	var c = [5]int{'A', 'B', 'C', 'D', 'E'} // byte
	fmt.Println(c)                          // [65 66 67 68 69]
	d := [...]int{1, 2, 3, 4, 5}            // 根据元素的个数，设置数组的大小
	fmt.Println(d)                          //[1 2 3 4 5]
	e := [5]int{4: 100}                     // [0 0 0 0 100]
	fmt.Println(e)
	f := [...]int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
	fmt.Println(f)

}
