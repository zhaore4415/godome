package main

import (
	"fmt"
	"reflect"
)
//检查类型和值
func inspect(v interface{}) {
	rv := reflect.ValueOf(v) //获取v 的值
	rt := reflect.TypeOf(v)  //获取v 的类型

	fmt.Printf("Type: %v\n", rt)
	fmt.Printf("Kind: %v\n", rt.Kind()) //返回变量的底层数据结构种类（Kind）。
	fmt.Printf("Value: %v\n", rv)
	fmt.Printf("Can set: %v\n", rv.CanSet()) //rv.CanSet() 判断这个 reflect.Value 是否可被修改。
}

func main() {
	var x int = 42
	inspect(x)
}
