// defer +recover 机制处理错误
// 参考：https://studygolang.com/articles/20196
// 使用defer + recover机制处理错误示例
package main

import (
	"fmt"
	"time"
)

func printNum() {
	for i := 1; i <= 10; i++ {

		fmt.Println(i)
	}
}

// 做除法操作
func devision() {
	// 使用defer + recover机制处理错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生错误:", err)
		}
	}()

	num1 := 10
	num2 := 0 //  这里故意设置为0，触发除零错误
	result := num1 / num2
	fmt.Println(result)
}

func main() {

	go printNum()
	go devision()

	// 模拟一些处理时间
	time.Sleep(2 * time.Second)

	fmt.Println("程序执行完毕")
}
