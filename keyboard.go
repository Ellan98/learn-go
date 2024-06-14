package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio包读取 1.先创建Reader对象：
func main() {
	// var x int
	// var y float64
	// fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	// fmt.Printf("x的数值：%d，y的数值：%f\n", x, y)
	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)
}

//打印：fmt.Print() (n int, err error )
//格式化打印：fmt.Printf() (n int, err error )
//打印后换行.Print() (n int, err error )
// 格式化打印占位符：

//             %v,原样输出
//             %T，打印类型
//             %t,bool类型
//             %s，字符串
//             %f，浮点
//             %d，10进制的整数
//             %b，2进制的整数
//             %o，8进制
//             %x，%X，16进制
//                 %x：0-9，a-f
//                 %X：0-9，A-F
//             %c，打印字符
//             %p，打印地址
//             。。。
