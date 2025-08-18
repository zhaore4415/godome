// repository/todo_repository.go
package repository

import (
	"database/sql"
	"web03/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

// CreateTodo 插入新 Todo
func (r *TodoRepository) CreateTodo(todo *models.Todo) error {
	query := `INSERT INTO todos (title, done,created_at) VALUES (?, ?, ?)`
	result, err := r.DB.Exec(query, todo.Title, todo.Done, todo.CreatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.ID = uint(id)
	return nil
}

// GetAllTodos 查询所有 Todos
func (r *TodoRepository) GetAllTodos() ([]models.Todo, error) {
	query := `SELECT id, title, done, created_at, updated_at FROM todos ORDER BY id`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, rows.Err()
}

// GetTodoByID 根据 ID 查询 Todo
func (r *TodoRepository) GetTodoByID(id uint) (*models.Todo, error) {
	query := `SELECT id, title, done, created_at, updated_at FROM todos WHERE id = ?`
	row := r.DB.QueryRow(query, id)

	var todo models.Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 返回 nil 表示未找到
		}
		return nil, err
	}
	return &todo, nil
}
