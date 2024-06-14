package main

import "fmt"

// import . "fmt" 可以省略前缀的包名 例如 Println("hello world")
// import f "fmt"  别名 例如 f.Println("hello world")
// import _ "github.com/ziutek/mymysql/godrv"  引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数

func main() {
	fmt.Println("hello world")
}
