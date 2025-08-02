package main

import "fmt"

func main() {

	//示例1
	var intChan chan int = make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-intChan)
	// }
	close(intChan) // 关闭管道 ,如果不关闭，下面循环取值会一直取，超过10就报deadlock错误
	//管理没有索引，直接用值接收
	for v := range intChan {
		fmt.Println("v:", v)
	}

}
