package main

import "fmt"

type Sayer interface {
	Say(word string)
}

type Persion struct {
	name string
}

// 方法接收者         方法签名 及其 参数
func (p Persion) Say(word string) {
	fmt.Printf("hello 我是%v,我说：%v", p.name, word)
}

func main() {

	p := Persion{
		name: "Ellan",
	}

	p.Say("destiny")

}
