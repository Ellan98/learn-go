package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//指针的使用
// &（取地址）和*（根据地址取值）。
// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
// 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
// 指针变量的值是指针地址。
// 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
func main() {
	// 指针类型初始化
	p1 := &Person{Name: "Alice", Age: 30}
	p3 := p1
	p2 := Person{Name: "Bob", Age: 25}
	p4 := p2
	p1.Name = "Ellan"
	p2.Name = "AAA"
	fmt.Printf("我是p1：%+v\n", p1)
	fmt.Printf("我是p2：%+v\n", p2)
	fmt.Printf("我是p3：%+v\n", p3)
	fmt.Printf("我是p4：%+v\n", p4)
}
