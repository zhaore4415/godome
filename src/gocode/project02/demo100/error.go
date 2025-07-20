package main

import "fmt"

func main() {
	testerr()
	fmt.Println("上面的除法操作执行成功")
	fmt.Println("正常执行下面的逻辑。。。")
}

func testerr() {
	// 选用defer + recover 来捕获错误：defer 后面加上匿名函数的调用   -- () 括号就是调用匿名函数
	defer func() {
		err := recover()
		//不为空就输出
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("err:", err)

		}
	}()

	num1 := 10
	num2 := 0

	result := num1 / num2
	fmt.Println(result)
}
