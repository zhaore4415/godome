package main

import "fmt"

func main() {

	// sum := 1
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// //死循环
	// for{
	// 	fmt.Printf("333")
	// }

	str := "Hello, world!你好，世界！"
	//fmt.Println(str)
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c \n", str[i])
	// }
	// 通过字节进行遍历，暂不支持中文
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])

	}
	//for rang
	for i, c := range str {
		fmt.Printf("%d %c \n", i, c)
	}
}
