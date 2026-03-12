package routes

import (
	"github.com/gin-gonic/gin"

	"task-service/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/tasks", handlers.GetTasks)
		v1.GET("/tasks/:id", handlers.GetTask)
		v1.POST("/tasks", handlers.CreateTask)
		v1.PUT("/tasks/:id", handlers.UpdateTask)
		v1.DELETE("/tasks/:id", handlers.DeleteTask)
	}

	return r
}
