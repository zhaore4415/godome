// repository/todo_repository.go
package repository_gorm

import (
	"web03/models"

	"gorm.io/gorm"
)

// 将 *sql.DB 改为 *gorm.DB
type TodoRepository struct {
	DB *gorm.DB // 使用 GORM 的 DB 实例
}

// NewTodoRepository 接收 *gorm.DB
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

// CreateTodo 插入新 Todo
func (r *TodoRepository) CreateTodo(todo *models.TodoGorm) error {
	// GORM 会自动处理 CreatedAt 和 UpdatedAt (如果用了 autoCreateTime 标签)
	// 不需要手动传入时间
	return r.DB.Create(todo).Error
	// .Error 会返回操作的错误，如果没有错误则为 nil
}

// GetAllTodos 查询所有 Todos
func (r *TodoRepository) GetAllTodos() ([]models.TodoGorm, error) {
	var todos []models.TodoGorm
	// GORM 的 Find 方法
	result := r.DB.Find(&todos)
	return todos, result.Error
}

// GetTodoByID 根据 ID 查询 Todo
func (r *TodoRepository) GetTodoByID(id uint) (*models.TodoGorm, error) {
	var todo models.TodoGorm
	// GORM 的 First 方法，根据主键查找
	result := r.DB.First(&todo, id) // id 作为主键值
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // 返回 nil 表示未找到
		}
		return nil, result.Error // 其他数据库错误
	}
	return &todo, nil
}
