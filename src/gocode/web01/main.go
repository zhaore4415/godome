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

	// 加载 HTML 模板（支持嵌套目录，通配符）D:\goproject\src\gocode\web01\templates\products\index.tmpl
	//r.LoadHTMLGlob("templates/**/*.tmpl") // 注意：部分系统仍不支持 **// 建议把模板放在 templates/ 目录下
	// 或者：r.LoadHTMLFiles("index.tmpl") 如果只有一个文件
	// 改成显式加载：
	// ✅ 使用 Glob 匹配所有嵌套目录下的 index.tmpl
	r.LoadHTMLGlob("templates/**/index.tmpl")

	// ✅ 打印调试信息（虽然不能直接访问 Templates，但我们可以信任 Gin）
	log.Println("✅ 已加载模板模式: templates/**/index.tmpl")

	data := PageVariables{
		Title:   "My Website，产品页面",
		Message: "Welcome to my website!",
		Items:   []string{"Item 1", "Item 2", "Item 3"},
	}
	u2 := User{
		Name:   "myzhaohuan",
		Gender: "男",
		Age:    18,
	}

	r.GET("/products/index", func(c *gin.Context) {
		// ✅ 使用 "index.tmpl" 作为模板名（因为文件名是 index.tmpl）
		c.HTML(200, "index.tmpl", gin.H{
			"Title":   data.Title,
			"Message": data.Message,
			"Items":   data.Items,
			"u2":      u2,
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"Title":   data.Title,
			"Message": data.Message,
			"Items":   data.Items,
			"u2":      u2,
		})
	})

	log.Println("🚀 服务器已启动，请访问: http://localhost:8081")
	err := r.Run(":8081")
	if err != nil {
		log.Fatal("❌ 启动失败:", err)
	}
}
