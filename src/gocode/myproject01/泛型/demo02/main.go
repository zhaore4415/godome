package main

import (
	"fmt"
)

// 类型约束:定义一个类型约束，仅允许int, int32, float32, float64类型
type Number interface {
	int | int32 | float32 | float64
}

// 定义一个泛型函数，用于计算两个数的和，且其类型必须满足Number约束
func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	// 使用Add函数进行不同类型的数值相加
	intResult := Add[int](10, 20)
	fmt.Println(intResult) // 输出: 30

	floatResult := Add[float64](5.5, 4.3)
	fmt.Println(floatResult) // 输出: 9.8

	fmt.Println(Add(1, 2))     // int  自动推断为int
	fmt.Println(Add(1.5, 2.7)) // float64   自动推断为float64
	// 下面这行会报错，因为string不在Number约束内
	// fmt.Println(Add("Hello", "World"))
}
