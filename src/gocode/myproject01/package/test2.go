package main

import "fmt"

//同一个包下面只能有一个mian 函数
func myMain() {
	// go 中每个文件都必须归属于一个包
	// 引入了 "fmt" 包，可以使用包里的函数
	fmt.Println("Hello34, go")
}
