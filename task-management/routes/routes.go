package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/KeerthanaThirugnanam/task-management/handlers"
	"github.com/KeerthanaThirugnanam/task-management/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", handlers.Login)
	router.POST("/tasks", middleware.AuthMiddleware(), handlers.CreateTask)
	router.GET("/tasks", middleware.AuthMiddleware(), handlers.GetTasks)
	router.GET("/ai/suggest", middleware.AuthMiddleware(), handlers.GetTaskSuggestions)
}
