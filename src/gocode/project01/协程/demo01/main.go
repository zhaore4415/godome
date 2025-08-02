package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond) // 简单等待，生产环境不推荐
		fmt.Println("Hello from Goroutine!")
	}
}

// main 函数是程序的入口点
func main() {
	// 启动一个 Goroutine 执行 sayHello 函数
	go sayHello()

	// 主 Goroutine 需要等待，否则程序可能在 sayHello 执行前就结束了

	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond) // 简单等待，生产环境不推荐
		fmt.Println("Hello from main!")
	}

}
