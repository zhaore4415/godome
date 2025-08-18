// main.go
package main

import (
	"log"
	"web03/database"
	_ "web03/docs" // docs 包 (由 swag 生成
	"web03/handlers"
	"web03/repository_gorm"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Sample API
// @version         1.0
// @description     This is a sample server for a Todo List API.
// @host      localhost:8082
// @BasePath  /api/v1
func main() {
	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer database.CloseDB()

	// // 创建 Repository  非gorm
	// todoRepo := repository.NewTodoRepository(database.DB)

	// // 创建 Handler
	// todoHandler := handlers.NewTodoHandler(todoRepo)

	// 2. 获取 GORM DB 实例 --使用gorm
	db := database.DB // 这里是 *gorm.DB
	// 3. 创建 Repository 实例 (传入 *gorm.DB)
	todoRepo := repository_gorm.NewTodoRepository(db)

	// 4. 创建 Handler 实例
	todoHandler := handlers.NewTodoHandler(todoRepo)

	r := gin.Default()

	// 设置路由组
	v1 := r.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.GET("", todoHandler.GetTodos)
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("/:id", todoHandler.GetTodo)
		}
	}

	// 添加 Swagger UI 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("服务器启动在 :8082")
	log.Println("访问 http://localhost:8082/swagger/index.html 查看 API 文档")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
