// models/todo.go
package models

import "time"

// Todo 模型
// @Description Todo 项目模型
type Todo struct {
	// ID of the todo item
	// in: path
	// required: true
	// example: 1
	ID uint `json:"id" example:"1"`
	// Title of the todo
	// required: true
	// min length: 1
	// max length: 50
	Title string `json:"title" binding:"required" example:"Learn Go"`
	// Whether the todo is completed
	Done      bool      `json:"done" example:"false"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Todo 模型
// @Description Todo 项目模型
type TodoGorm struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Title     string    `json:"title" binding:"required,min=1,max=50" gorm:"not null" example:"Learn Go"`
	Done      bool      `json:"done" gorm:"default:false" example:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"` // GORM 会自动处理
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"` // GORM 会自动处理
}

// ErrorResponse 模型
// @Description 通用错误响应
type ErrorResponse struct {
	Error string `json:"error" example:"Something went wrong"`
}
