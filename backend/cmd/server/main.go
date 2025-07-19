package main

import (
	"sandbox/internal/presentation/http"

	"github.com/gin-gonic/gin"
)

// @title Todo API
// @version 1.0
// @description This is a sample server for a todo application.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	router := http.NewRouter()
	router.SetupRoutes(r)
	r.Run(":8080")
}
