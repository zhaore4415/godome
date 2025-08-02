package main

import (
	"fmt"
	"time"
)

// 使用select实现多路复用的管道示例
func main() {
	var intChan1 = make(chan int)
	var intChan2 = make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second) // 模拟一些处理时间
			intChan1 <- i
		}
		close(intChan1)
	}()

	go func() {
		for i := 5; i < 10; i++ {
			intChan2 <- i
		}
		close(intChan2)
	}()

		for{
			select {
			case num, ok := <-intChan1:
				if ok {
					fmt.Println("从intChan1中读取到的值是:", num)
				}
			case num, ok := <-intChan2:
				if ok {
					fmt.Println("从intChan2中读取到的值是:", num)
				}
			default:
				fmt.Println("没有可读数据")
			}
		}
}
