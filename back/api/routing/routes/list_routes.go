package routes

import (
	// middlewares "github.com/saegus/test-technique-romain-chenard/internal/middleware"

	listCtrl "github.com/saegus/test-technique-romain-chenard/api/controllers"
	"github.com/saegus/test-technique-romain-chenard/api/middlewares"

	"github.com/gin-gonic/gin"
)

func ListRoutes(router *gin.Engine) {

	listController := listCtrl.New()
	guestGroup := router.Group("/todo/list")
	{
		guestGroup.POST("", middlewares.IsAuth(false), listController.CreateList)
		guestGroup.GET("", middlewares.IsAuth(false), listController.GetLists)
		guestGroup.DELETE("/:listId", middlewares.IsAuth(false), listController.DeleteList)
		guestGroup.PUT("/:listId", middlewares.IsAuth(false), listController.UpdateList)
	}	
}
