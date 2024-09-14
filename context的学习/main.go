/*
 * @Date: 2024-08-08 13:42:43
 * @LastEditTime: 2024-09-11 16:51:32
 * @FilePath: \context的学习\main.go
 * @description: 注释
 */
// golang context  是 golang 中的经典工具，主要在异步场景中用于实现并发协调以及对 goroutine 的生命周期控制. 除此之外，context 还兼有一定的数据存储能力.
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// context.Context 是一种interface 类型 共有4种 核心 API
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

// 标准 error
// Canceled是取消上下文时[Context.Err]返回的错误
var Canceled = errors.New("context canceled")

// DeadlineExceded是[Context.Err]在上下文发生错误时返回的错误
// 截止日期过去了。
var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string   { return "context deadline exceeded" }
func (deadlineExceededError) Timeout() bool   { return true }
func (deadlineExceededError) Temporary() bool { return true }

//2.emptyCtx
// 2.1 emptyCtx 是一个空的 context，本质上类型为一个整型；

// 2.2Deadline 方法会返回一个公元元年时间以及 false 的 flag，标识当前 context 不存在过期时间；

// 2.3Done 方法返回一个 nil 值，用户无论往 nil 中写入或者读取数据，均会陷入阻塞；

// 2.4Err 方法返回的错误永远为 nil；

// 2.5Value 方法返回的 value 同样永远为 nil.
//========================================
// context.Background() & context.TODO() 均返回emptyCtx类型的一个实例
//========================================
//  cancelCtx 数据结构 为什么 会带 互斥锁？
//内置 一个 context 作为 父 ,所以cancelCtx一定 是 某个 context的子context
//内置一把 互斥锁 用以协调并发场景下的资源获取
// done 实际类型为 chan struct{} 用于以反映cancelCtx生命周期的通道
// children 一个set 指向cancelCtx下的所有子context
//err 记录 当前cancelCtx的错误 必然为某一个context的子context ?为什么这么说
