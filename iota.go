// iota，特殊常量，可以认为是一个可以被编译器修改的常量
package main

import "fmt"

// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
func main() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数 中断iota自增，则必须显式恢复。且后续自增值按行序递增 自增默认是int类型，可以自行进行显示指定类型
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

// 数字常量不会分配存储空间，无须像变量那样通过内存寻址来取值，因此无法获取地址
