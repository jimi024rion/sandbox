package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	todoHandler "sandbox/internal/presentation/http/todos"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (ro *Router) SetupRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		todos := api.Group("/todos")
		{
			todos.GET("", ro.GetTodos)
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("/:id", ro.GetTodo)
			todos.PUT("/:id", ro.UpdateTodo)
			todos.DELETE("/:id", ro.DeleteTodo)
		}
	}
}

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /todos [get]
func (ro *Router) GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTodos"})
}

// GetTodo godoc
// @Summary Get a todo by ID
// @Description Get a todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {string} string "ok"
// @Router /todos/{id} [get]
func (ro *Router) GetTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTodo"})
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update a todo
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body entity.Todo true "Todo object"
// @Success 200 {string} string "updated"
// @Router /todos/{id} [put]
func (ro *Router) UpdateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateTodo"})
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {string} string "deleted"
// @Router /todos/{id} [delete]
func (ro *Router) DeleteTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteTodo"})
}
