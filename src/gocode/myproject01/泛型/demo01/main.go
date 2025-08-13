package main

import "fmt"

// 类型参数:定义一个泛型函数，接受任意类型的参数
func Print[T any](item T) { //any 是 interface{} 的别名,它表示“任意类型”。
	fmt.Println(item)
}

//// 这两行是等价的（语义上）：
// func Print[T any](item T)
// func Print[T interface{}](item T)

//只允许可比较的类型（如 int, string, struct 等）
func Equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	// 调用Print函数，传入不同类型的数据
	Print(42)      // int
	Print("Hello") // string
	Print(5.56)    // string
}
