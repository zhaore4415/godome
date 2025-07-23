//定义数据类型
package main

import "fmt"

func main(){
	  type myint int //定义了一个myint类型，底层类型为int
    var b1 myint = 20
    fmt.Printf("b1:%T\n", b1) //打印b1的类型
    fmt.Println(b1)           //打印b1的值
    // test(b1)                  //调用test函数，传入参数b的int值
    var b2 int = 30
    b2 = int(b1)
    fmt.Println(b2) //打印b2的值
}