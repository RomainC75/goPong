package routes

import (
	socketCtrl "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	socketController := socketCtrl.New()
	// guestGroup := router.Group("/ws")
	// {
	// 	guestGroup.GET("/", socketController.Socket)
		
	// }
	// !!!!!!!!!!!!!!!! WARNING !!!!!!!!!!!!!!!!!!!!
	router.GET("/ws", socketController.Socket)
}
