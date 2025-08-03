package main

//修改值（需要传入指针）
import (
	"fmt"
	"reflect"
)

func setValue(v interface{}) {
	rv := reflect.ValueOf(v)

	// 如果传入的是指针，需要获取指针指向的元素
	//检查 rv 是否是一个指针类型（Kind 是 Ptr）。
	// 如果是，则调用 rv.Elem() 获取指针所指向的目标对象的反射值。
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem() //返回 i 本身的反射值（即 int 类型的 10）
	}
	//rv.CanSet() 判断这个反射值是否可被修改。
	if rv.CanSet() {
		switch rv.Kind() {
		case reflect.Int: //如果是整型（int, int32, int64 等的 Kind 都是 int），就设置为 100
			rv.SetInt(100)
			// case reflect.String:/如果是字符串类型，就改为 "modified"
			rv.SetString("modified")
		}
	}
}

func main() {
	var i int = 10
	var s string = "hello"

	setValue(&i)
	setValue(&s)

	fmt.Println(i, s) // 100 modified
}
