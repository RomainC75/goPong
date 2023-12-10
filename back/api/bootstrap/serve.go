package bootstrap

import (
	"github.com/saegus/test-technique-romain-chenard/config"
	"github.com/saegus/test-technique-romain-chenard/data/database"

	Routing "github.com/saegus/test-technique-romain-chenard/api/routing"
)

func Serve() {
	config.Set()

	database.Connect()

	Routing.Init()

	Routing.RegisterRoutes()

	Routing.Serve()
}
