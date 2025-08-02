package main

import "fmt"

//声明一个只读，和只写的管道示例

func main() {
	// var intChan chan int = make(chan int, 3)

	var intChan2 chan<- int //  <-  在后面，代表往 chan写
	intChan2 = make(chan int, 3)
	intChan2 <- 20
	fmt.Println("intChan2", intChan2, "intChan2长度:", len(intChan2), "intChan2容量:", cap(intChan2))

	var intChan3 <-chan int //  <-  在前面 代表从 chan读
	if intChan3 != nil {
		num := <-intChan3
		fmt.Println("从intChan3中读取到的值是:", num)
	}

}
