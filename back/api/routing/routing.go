package routing

import (
	"github.com/saegus/test-technique-romain-chenard/internal/providers/routes"

	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())
}
