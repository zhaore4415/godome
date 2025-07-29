package main

import (
	"fmt"
	"strconv"
)

// 字符转换
func main() {
	// 字符转换
	var char byte = 'A'                      // 字符 'A' 的 ASCII 码是 65
	fmt.Println("ASCII value of 'A':", char) // 输出: ASCII value of 'A': 65

	var num int = 66                                   // 整数 66 对应的字符是 'B'
	char = byte(num)                                   // 将整数转换为字符
	fmt.Println("Character for ASCII value 66:", char) // 输出: Character for ASCII value 66: B

	// 字符串转换
	str := "Hello, World!"
	fmt.Println("String:", str) // 输出: String: Hello, World!

	// 字符串转int //这里要做判断 转换失败怎么办
	var num2 int
	num2, err := fmt.Sscanf("中123", "%d", &num2) // Sscanf 会从字符串开头开始解析，直到遇到非数字字符。最终 会解析出整数部分 123

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Integer value of '123':", num2)
	}

	str2 := "123中"
	num3, err1 := strconv.Atoi(str2) //strconv.Atoi 要求整个字符串必须是合法整数，否则返回错误。
	if err1 != nil {
		fmt.Println("转换失败:", err1)
	} else {
		fmt.Println("Integer value of '123':", num3)
	}

	// int转字符串
	str3 := fmt.Sprintf("%d", num3)
	fmt.Println("String value of 123:", str3) // 输出: String value of 123: 123

}
