// handlers/todo_handler.go
package handlers

import (
	"net/http"
	"strconv"
	"web03/models"
	"web03/repository_gorm"

	"github.com/gin-gonic/gin"
)

// 将全局变量 todos 移除，使用 repository

// 定义一个结构体来持有依赖
type TodoHandler struct {
	Repo *repository_gorm.TodoRepository
}

// NewTodoHandler 创建新的 Handler
func NewTodoHandler(repo *repository_gorm.TodoRepository) *TodoHandler {
	return &TodoHandler{Repo: repo}
}

// @Summary 创建一个新的 Todo 项目
// @Description 创建一个新的待办事项
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo 项目（无需传 ID）"
// @Success 201 {object} models.Todo
// @Failure 400 {object} models.ErrorResponse
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.TodoGorm
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "无效的请求数据: " + err.Error(),
		})
		return
	}

	// Done 字段由客户端决定或默认 false
	// todo.Done = false // 如果需要强制未完成，可以取消注释

	if err := h.Repo.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodos 获取所有 Todos
// @Summary 获取所有待办事项
// @Description 获取所有待办事项的列表
// @Tags todos
// @Produce json
// @Success 200 {array} models.Todo
// @Failure 500 {object} models.ErrorResponse
// @Router /todos [get]
func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.Repo.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "查询失败: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GetTodo 根据 ID 获取 Todo
// @Summary 根据 ID 获取单个待办事项
// @Description 根据提供的 ID 获取一个具体的待办事项
// @Tags todos
// @Produce json
// @Param id path int true "待办事项 ID"
// @Success 200 {object} models.Todo
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "无效的 ID"})
		return
	}

	todo, err := h.Repo.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "查询失败: " + err.Error(),
		})
		return
	}
	if todo == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "未找到"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
