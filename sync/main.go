package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m     sync.Mutex //互斥锁
	count int        //共享资源
	w     sync.WaitGroup
}

func (c *Counter) Increment() {
	c.count++
}
func (c *Counter) Decrement() {
	c.count--
}

func main() {

	counter := &Counter{}

	for i := 0; i < 10; i++ {
		counter.w.Add(1)
		go func() {
			defer counter.w.Done()
			counter.Increment()
			counter.Decrement()
		}()
	}

	counter.w.Wait()

	fmt.Println("最终数值：", counter.count)

}
