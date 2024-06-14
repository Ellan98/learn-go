package main

import "fmt"

// for init; condition; post { }  结构

func main() {
	for i := 1; i <= 10; {
		i = i + 1
		fmt.Printf(" %d", i)
	}

}

// 所有的三个组成部分，即初始化、条件和post都是可选的。
// for condition { }
// 效果与while相似
// for { }
