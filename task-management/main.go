package main

import (
	"github.com/gin-gonic/gin"
	"github.com/KeerthanaThirugnanam/task-management/database"
	"github.com/KeerthanaThirugnanam/task-management/routes"
)

func main() {
	database.ConnectDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
