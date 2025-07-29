package main

//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
const (
	a = iota
	b
	c
)

func main() {

	println(a, b, c)
}
