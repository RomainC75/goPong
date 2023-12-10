package routes

import (
	"github.com/saegus/test-technique-romain-chenard/api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(middlewares.CORSMiddleware())
	UserRoutes(router)
	ListRoutes(router)
	TaskRoutes(router)
	SocketRoutes(router)

}
