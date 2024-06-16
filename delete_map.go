package main

import "fmt"

//delete(map,key)函数用于 函数集合的元素 参数为map 和对应的key 删除的函数 不返回任何值
func main() {

	data_map := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	delete(data_map, "key2")
	fmt.Println("处理后的map", data_map)

}
