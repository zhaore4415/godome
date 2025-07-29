//func
//函数名(形参列表)(返回值类型列表){
//执行语句..
//return +返回值列表
//}

package main

import "fmt"

func calc(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	var a int = 10
	dd := calc(a, 20)
	fmt.Println(dd)
}
