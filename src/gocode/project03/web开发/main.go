package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, 你好！这是你的第一个 Go Web 服务！\n")
	fmt.Fprintf(w, "请求路径: %s\n", r.URL.Path)
}

func main() {
	// 注册路由和处理器
	http.HandleFunc("/", helloHandler)

	// 启动服务器，监听 8080 端口
	fmt.Println("服务器启动在 http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}