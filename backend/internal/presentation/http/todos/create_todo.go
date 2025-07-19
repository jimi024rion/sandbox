package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sandbox/internal/domain/entity/todos"
)

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body todos.Todo true "Todo object"
// @Success 201 {object} todos.Todo
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var todo todos.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}
