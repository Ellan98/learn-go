package main

import "fmt"

// interface
type Sayer interface {
	Say(mover string)
}

type Person struct {
	name string
}

//方法要写在 main外
func (p Person) Say(mover string) {
	fmt.Printf("名字: %+v,say:%+v\n", p.name, mover)
}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	//实例化结构体
	p := Person{
		name: "张三",
	}
	//调用并传入参数
	p.Say("hello golang")

	// 结构体 与 接口 方法签名一致 才可以 将该结构体的实例可以被赋值给接口类型的变量，从而通过接口类型来调用这些方法。
	var s Sayer = p
	s.Say("again implement ")

	d := dryer{}

	h := haier{
		dryer: d,
	}

	h.dry()

}
