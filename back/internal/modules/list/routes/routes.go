package routes

import (
	// middlewares "github.com/saegus/test-technique-romain-chenard/internal/middleware"

	"github.com/saegus/test-technique-romain-chenard/internal/middlewares"
	listCtrl "github.com/saegus/test-technique-romain-chenard/internal/modules/list/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	listController := listCtrl.New()
	guestGroup := router.Group("/todo/list")
	{
		guestGroup.POST("", middlewares.IsAuth(), listController.CreateList)
		guestGroup.GET("", middlewares.IsAuth(), listController.GetLists)
		guestGroup.DELETE("/:listId", middlewares.IsAuth(), listController.DeleteList)
		guestGroup.PUT("/:listId", middlewares.IsAuth(), listController.UpdateList)
	}	
}
