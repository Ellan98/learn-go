// /*
//   - @Date: 2024-07-24 16:47:25
//   - @LastEditTime: 2024-07-31 17:42:55
//   - @FilePath: \goroutine\main.go
//   - @description: 注释
//     */
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var initNumber int32

func main() { //开启一个主 goroutine 去执行 main函数

	wg.Add(1) // 计数牌
	initNumber = 10
	go editNumber(11) // 开启 goroutine 去执行
	go editNumber(12) //

	wg.Wait() //等待 goroutine 执行结束

	fmt.Printf("当前数值为%v\n", initNumber)
	// time.Sleep(time.Second) 不建议使用

}

func editNumber(num int32) {

	defer wg.Done() //  阻塞    通知wg 计数器-1 等待计数器 为 0
	//如果不加锁 就会导致 数据中途被篡改 也就是 进入 数据竞争态
	mutex.Lock()

	initNumber = num + initNumber
	fmt.Printf("参数为%v,此时initNumber:%v\n", num, initNumber)

	mutex.Unlock()

}

// ==================================================== 并发模拟
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var (
// 	wg      sync.WaitGroup
// 	counter int32
// 	mu      sync.Mutex
// )

// func main() {
// 	wg.Add(2) // 等待两个 goroutine 完成

// 	go increment()
// 	go increment()

// 	wg.Wait() // 等待所有 goroutine 完成
// 	fmt.Printf("最终计数器值为: %v\n", counter)
// }

// func increment() {
// 	defer wg.Done()

// 	for i := 0; i < 1000; i++ {
// 		mu.Lock()
// 		counter++ // 竞态条件发生在这里
// 		mu.Unlock()
// 	}
// }
