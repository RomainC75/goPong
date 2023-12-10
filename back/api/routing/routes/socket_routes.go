package routes

import (
	"github.com/gin-gonic/gin"
	socketCtrl "github.com/saegus/test-technique-romain-chenard/api/controllers"
	"github.com/saegus/test-technique-romain-chenard/api/middlewares"
)

func SocketRoutes(router *gin.Engine) {

	socketController := socketCtrl.New()
	// guestGroup := router.Group("/ws")
	// {
	// 	guestGroup.GET("/", socketController.Socket)
		
	// }
	// !!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!
	
	router.GET("/ws", middlewares.IsAuth(true),socketController.Socket)
}
