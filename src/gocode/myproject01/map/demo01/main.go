package main

import "fmt"

func main() {
	//map 是无序的键值对集合
	//map的键可以是任意类型，值也可以是任意类型
	//map的几种创建方式
	//1. 使用字面量创建map
	myMap2 := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 15,
	}
	for key, value := range myMap2 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	fmt.Printf("----------\n")
	//2. 使用make函数创建一个空的map
	myMap3 := make(map[string]int)
	myMap3["apple"] = 5
	myMap3["banana"] = 10
	myMap3["orange"] = 15
	for key, value := range myMap3 {
		fmt.Printf("Key: %s, Value: %d\n", key, value) //%s 表示字符串，%d表示整数
	}
	fmt.Printf("----------\n")
	//3. 使用make函数创建一个指定大小的map
	myMap4 := make(map[string]int, 10) // 创建一个空的map，		// 但不指定初始键值对
	myMap4["apple"] = 5
	myMap4["banana"] = 10
	myMap4["orange"] = 15
	for key, value := range myMap4 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	//

}
