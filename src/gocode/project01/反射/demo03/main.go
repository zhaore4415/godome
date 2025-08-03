package main

import (
	"fmt"
	"reflect"
)

// 调用方法
type Calculator struct{}

func (c Calculator) Add(a, b int) int {
	return a + b
}

func (c Calculator) Multiply(a, b int) int {
	return a * b
}

func main() {
	calc := Calculator{}
	rv := reflect.ValueOf(calc)

	// 调用 Add 方法，根据方法名称
	addMethod := rv.MethodByName("Add")

	// 反射调用方法时，参数必须是 []reflect.Value 类型。
	// 所以我们要把普通值 10 和 20 包装成 reflect.Value。
	// 顺序必须和方法签名一致：Add(a, b int) → 第一个参数是 a，第二个是 b
	args := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	}
	result := addMethod.Call(args)
	fmt.Println("10 + 20 =", result[0].Int()) // 30

	// 调用 Multiply 方法 ，根据方法索引
	addMethod2 := rv.Method(1) // 获取第二个方法，即 Multiply ，顺序是根据ASCII码排序的，a,b,...，索引0,1,2,...

	args2 := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	}
	result2 := addMethod2.Call(args2)
	fmt.Println("10 * 20 =", result2[0].Int()) // 200
}
