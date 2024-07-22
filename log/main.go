/*
 * @Author:
 * @Date: 2024-07-22 19:44:33
 * @LastEditors:
 * @LastEditTime: 2024-07-22 20:01:34
 * @Description:
 */
package main

import (
	"log"
	"os"
)

// log 日志库
func main() {
	//日志的配置
	//设置日志前缀
	log.SetPrefix("Test:")
	//设置格式
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	//设置日志的输出位置
	f, err := os.OpenFile("./cli.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755) //如果文件不存咋 则 创建 只写入
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)
	//基本用法 带有时间
	log.Print("普通打印")
	log.Printf("格式化打印%s\n", "format")
	log.Println("打印并换行")
	// log 其它使用
	//报错 并显示 错误信息后 退出 执行
	// log.Panicln("panicln用法")
	//打印日志后退出
	// log.Fatalln("fatalln使用")
}
