/*
 * @Date: 2024-07-04 13:43:56
 * @LastEditTime: 2024-07-04 13:54:32
 * @FilePath: \simple_use_gin\func\main.go
 * @description: 注释
 */
package main

import "fmt"

//格式
// func 函数名(参数)(返回值){
// 	函数体
// }

//可变参数通常要作为函数的最后一个参数
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 函数如果有多个返回值时必须用()将所有返回值包裹起来。
// func calc(x, y int) (int, int) {
// 	sum := x + y
// 	sub := x - y
// 	return sum, sub
// }

// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
// func someFunc(x string) []int {
// 	if x == "" {
// 		return nil // 没必要返回[]int{}
// 	}

// }

// 定义函数类型
type calculation func(int, int) int

// 定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
// 凡是满足这个条件的函数都是calculation类型的函数

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// add和sub都能赋值给calculation类型的变量。
// var c calculation
// c = add

//匿名函数
// func(参数)(返回值){
// 	函数体
// }
// 匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

func main() {

	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	funcA()
	funcB()
	funcC()

}

//内置函数
// close 主要用来关闭channel
// len 用来求长度，比如string、array、slice、map、channel
// new 用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
// make 用来分配内存，主要用来分配引用类型，比如chan、map、slice
// append 用来追加元素到数组、slice中
// panic和recover 用来做错误处理

// panic可以在任何地方引发，但recover只有在defer调用的函数中有效。

func funcA() {
	fmt.Println("func A")
}

// 程序运行期间funcB中引发了panic导致程序崩溃，异常退出了。这个时候就可以通过recover将程序恢复回来，继续往后执行。

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

// recover()必须搭配defer使用。
// defer一定要在可能引发panic的语句之前定义。
