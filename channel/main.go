package main

import "fmt"

// 声明 chan 后必须使用 make 来进行初始化
// 通道 有 发送 接收 和 关闭 三种操作
// 通道是可以 被 垃圾机制回收的

func main() {
	// 声明通道变量并初始化
	ch1 := make(chan int)

	// 启动一个 goroutine 发送数据到通道
	go func() {
		ch1 <- 10
	}()

	// 从通道接收数据
	x := <-ch1
	fmt.Println("x:", x)

	// 关闭通道
	close(ch1)
}
