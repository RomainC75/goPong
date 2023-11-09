package routes

import (
	// middlewares "github.com/saegus/test-technique-romain-chenard/internal/middleware"

	"github.com/saegus/test-technique-romain-chenard/internal/middlewares"
	taskCtrl "github.com/saegus/test-technique-romain-chenard/internal/modules/task/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	taskController := taskCtrl.New()
	guestGroup := router.Group("/todo/task")
	{
		guestGroup.PUT("/toggle/:taskId", middlewares.IsAuth(), taskController.ToogleTask)
		guestGroup.POST("/:listId", middlewares.IsAuth(), taskController.CreateTask)
		guestGroup.GET("/:listId", middlewares.IsAuth(), taskController.GetTasks)
		guestGroup.PUT("/:taskId", middlewares.IsAuth(), taskController.UpdateTask)
		guestGroup.DELETE("/:taskId", middlewares.IsAuth(), taskController.DeleteTask)
	}	
}
