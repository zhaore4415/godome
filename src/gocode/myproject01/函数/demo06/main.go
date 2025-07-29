package main

import "fmt"

//Go 语言递归函数
func main() {
	// 递归函数示例：计算阶乘
	var factorial func(n int) int
	factorial = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}

	// 测试递归函数
	num := 5
	result := factorial(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
	//内部计算逻辑 factorial(5) = 5 * factorial(4)
	//            							= 5 * 4 * factorial(3)
	//            							...
	//            							= 5 * 4 * 3 * 2 * 1 = 120

	// 递归函数示例：斐波那契数列
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	// 测试斐波那契函数
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
	}
	//内部逻辑 	fib(0) = 0
	// 					fib(1) = 1
	// 					fib(n) = fib(n-1) + fib(n-2) （递归定义）
}
