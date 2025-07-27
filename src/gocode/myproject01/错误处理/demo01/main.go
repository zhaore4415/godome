package main

import "fmt"

// func main() {
// 	testerr()
// 	fmt.Println("上面的除法操作执行成功")
// 	fmt.Println("正常执行下面的逻辑。。。")
// }

// func testerr() {

// 	num1 := 10
// 	num2 := 0

// 	result := num1 / num2
// 	fmt.Println(result)
// }

//defer + recover 处理错误
func main() {
	testerr()
	fmt.Println("上面的除法操作执行成功")
	fmt.Println("正常执行下面的逻辑。。。")
}

func testerr() {
	//defer + recover +匿名函数
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到错误：", err)
		}

	}()

	num1 := 10
	num2 := 0

	result := num1 / num2
	fmt.Println(result)
}
