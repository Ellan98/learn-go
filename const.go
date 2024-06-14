package main

import "fmt"

// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
// 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
// 不曾使用的常量，在编译的时候，是不会报错的
// 显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，变量是可以是不同的类型值

// const identifier [type] = value
// 显式类型定义： const b string = "abc"
// 隐式类型定义： const b = "abc"

// 常量可以作为枚举，常量组
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

// 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
const (
	x uint16 = 16
	y
	s = "abc"
	z
)
const b string = "abc"

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
	fmt.Printf("%T,%v\n", y, y)
	fmt.Printf("%T,%v\n", z, z)
}
