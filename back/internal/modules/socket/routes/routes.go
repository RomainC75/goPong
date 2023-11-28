package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saegus/test-technique-romain-chenard/internal/middlewares"
	socketCtrl "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/controllers"
)

func Routes(router *gin.Engine) {

	socketController := socketCtrl.New()
	// guestGroup := router.Group("/ws")
	// {
	// 	guestGroup.GET("/", socketController.Socket)
		
	// }
	// !!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!
	
	router.GET("/ws", middlewares.IsAuth(true),socketController.Socket)
}
