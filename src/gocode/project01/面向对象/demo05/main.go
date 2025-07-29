package main

import "fmt"

type integer int

//结构体是值传递
func (i integer) print() {
	i = 10
	fmt.Println("1i=", i) // i= 10

}

//如果想改变结构体的值，可以用指针，变成引用传递
func (i *integer) print2() {
	*i = 10
	fmt.Println("2i=", *i) // i= 10

}
func main() {
	var i integer = 20
	i.print()
	i.print2()

	fmt.Println("i=", i) // i= 20
}
