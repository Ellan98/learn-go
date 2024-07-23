package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() { //开启一个主 goroutine 去执行 main函数
	wg.Add(1)      // 计数牌
	go printText() // 开启 一个 独立 的goroutine 去执行
	fmt.Println("go 之前打印")
	wg.Wait() //等待 goroutine 执行结束
	// time.Sleep(time.Second) 不建议使用

}

func printText() {
	fmt.Println("执行")
	wg.Done() //  阻塞    通知wg 计数器-1 等待计数器 为 0

}
