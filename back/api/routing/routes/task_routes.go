package routes

import (
	// middlewares "github.com/saegus/test-technique-romain-chenard/internal/middleware"

	Controllers "github.com/saegus/test-technique-romain-chenard/api/controllers"
	"github.com/saegus/test-technique-romain-chenard/api/middlewares"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {

	taskController := Controllers.NewTaskCtrl()
	guestGroup := router.Group("/todo/task")
	{
		guestGroup.PUT("/toggle/:taskId", middlewares.IsAuth(false), taskController.ToogleTask)
		guestGroup.POST("/:listId", middlewares.IsAuth(false), taskController.CreateTask)
		guestGroup.GET("/:listId", middlewares.IsAuth(false), taskController.GetTasks)
		guestGroup.PUT("/:taskId", middlewares.IsAuth(false), taskController.UpdateTask)
		guestGroup.DELETE("/:taskId", middlewares.IsAuth(false), taskController.DeleteTask)
	}	
}
