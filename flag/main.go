/*
 * @Author:
 * @Date: 2024-07-23 19:23:17
 * @LastEditors:
 * @LastEditTime: 2024-07-23 19:32:23
 * @Description:
 */
package main

import (
	"flag"
	"fmt"
)

// flag 命令行参数 解析 标准库

func main() {
	//命令参数名称   默认值   解释
	name := flag.String("name", "Ellan", "姓名")
	age := flag.Int("age", 12, "年龄")
	//不需要像上面那样赋值 而是直接绑定到 gender 变量上
	var gender string
	flag.StringVar(&gender, "gennder", "", "性别")
	// 将参数绑定到 结构体
	//

	// flag.Var()  待了解
	//
	flag.Parse()
	// go run main.go -name "alice" -age 18
	fmt.Printf("名字%v,年龄：%v,性别：%v", *name, *age, gender)

}
