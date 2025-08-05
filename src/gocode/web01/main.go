package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type PageVariables struct {
	Title   string
	Message string
	Items   []string
}

type User struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	// 创建默认的 Gin 路由
	r := gin.Default()

	// 加载 HTML 模板（支持嵌套目录，通配符）
	r.LoadHTMLGlob("templates/*") // 建议把模板放在 templates/ 目录下
	// 或者：r.LoadHTMLFiles("index.tmpl") 如果只有一个文件

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		data := PageVariables{
			Title:   "My Website",
			Message: "Welcome to my website!",
			Items:   []string{"Item 1", "Item 2", "Item 3"},
		}
		u2 := User{
			Name:   "myzhaohuan",
			Gender: "男",
			Age:    18,
		}

		// 传递多个变量到模板
		c.HTML(200, "index.tmpl", gin.H{
			"Title":   data.Title,
			"Message": data.Message,
			"Items":   data.Items,
			"u2":      u2,
		})

		// 使用 HTML 方法渲染模板并返回
		c.HTML(200, "index.tmpl", data)

	})
	// 如果有多个模板文件，可以使用下面的方式指定模板
	// c.HTML(200, "other.tmpl", data)
	// 启动服务器
	log.Println("服务器已启动，请访问: http://localhost:8081")
	err := r.Run(":8081") // 默认监听 8080 端口
	if err != nil {
		log.Fatal("启动失败:", err)
	}
}
