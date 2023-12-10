package routes

import (
	// middlewares "github.com/saegus/test-technique-romain-chenard/internal/middleware"

	Controllers "github.com/saegus/test-technique-romain-chenard/api/controllers"
	"github.com/saegus/test-technique-romain-chenard/api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	userController := Controllers.NewAuthCtrl()
	guestGroup := router.Group("/auth")
	{
		guestGroup.POST("/signup", userController.HandleSignup)
		guestGroup.POST("/signin", userController.HandleSignin)
		guestGroup.GET("/verify", middlewares.IsAuth(false), userController.Verify)
	}	
}
