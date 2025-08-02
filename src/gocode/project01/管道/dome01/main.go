package main

import "fmt"

func main() {

	//示例1
	// var intChan chan int = make(chan int, 3)

	// //管道是引用类型
	// fmt.Println(intChan)

	// // 向通道中写入数据
	// intChan <- 100
	// num := 20
	// intChan <- num
	// intChan <- 30
	// // intChan<-40  长度只3所以不能放四个

	// fmt.Println("intChan:", intChan, "intChan长度:", len(intChan), "intChan容量:", cap(intChan))
	// // 向通道中取出数据
	// data1 := <-intChan
	// data2 := <-intChan

	// //如果再取，就会报错，因为只塞进去两个值
	// data3 := <-intChan
	// fmt.Println("data1:", data1, "data2:", data2, "data3:", data3)

	// 示例2
	// 关闭管道
	var strChan chan string = make(chan string, 3)
	strChan <- "Hello"
	strChan <- "World"
	fmt.Println("strChan:", strChan, "strChan长度:", len(strChan), "strChan容量:", cap(strChan))
	str1 := <-strChan
	str2 := <-strChan
	close(strChan) //关闭管道后，不能再向管道写入数据，但可以读取数据
	fmt.Println("str1:", str1, "str2:", str2)
	//str3 := <-strChan
	//fmt.Println("str3:", str3)     // 读取关闭后的管道会返回零值
	strVal := <-strChan            //此时无法再向管道写入数据，但可以读数据
	fmt.Println("strVal:", strVal) // 读取关闭后的管道会返回零值

}
