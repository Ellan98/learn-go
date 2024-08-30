/*
 * @Date: 2024-08-08 13:42:43
 * @LastEditTime: 2024-08-30 17:13:16
 * @FilePath: \context的学习\main.go
 * @description: 注释
 */
// golang context  是 golang 中的经典工具，主要在异步场景中用于实现并发协调以及对 goroutine 的生命周期控制. 除此之外，context 还兼有一定的数据存储能力.
package main

import (
	"context"
	"fmt"
	"time"
)

// context.Context 是一种interface 类型 共有4种API
// 1.Deadline() (time.Time,bool)  返回Ctx过期时间
// 2.Done() <- chan struct{} 返回用以标识ctx是否结束的chan
// 3.Error() error  返回ctx的错误
// 4.Value(key any) any 返回ctx存放的对应于key的value

type ContextUnit struct {
	context context.Context
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {

		for {
			select {
			case <-ctx.Done():
				fmt.Println("迎来了生命周期的终点")
				return
			default:
				fmt.Println("持续执行中")
			}
		}

	}(ctx)

	fmt.Println("ctx", ctx)

	time.Sleep(8 * time.Second)
	cancel()
	time.Sleep(9 * time.Second)
}
