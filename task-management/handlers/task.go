package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/KeerthanaThirugnanam/task-management/database"
	"github.com/KeerthanaThirugnanam/task-management/models"
)

var taskCollection = database.GetCollection("tasks")

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task.ID = primitive.NewObjectID()
	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}
