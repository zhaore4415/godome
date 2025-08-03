package main
//结构体反射
import (
	"fmt"
	"reflect"
)
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}

func main() {
    p := Person{Name: "Alice", Age: 25, City: "Beijing"}
    rv := reflect.ValueOf(p)
    rt := reflect.TypeOf(p)
    
    // 遍历结构体字段
    for i := 0; i < rv.NumField(); i++ {
        field := rt.Field(i)
        value := rv.Field(i)
        
        fmt.Printf("Field: %s, Type: %s, Value: %v", 
            field.Name, field.Type, value)
        
        // 获取 tag
        if jsonTag := field.Tag.Get("json"); jsonTag != "" {
            fmt.Printf(", JSON tag: %s", jsonTag)
        }
        fmt.Println()
    }
}