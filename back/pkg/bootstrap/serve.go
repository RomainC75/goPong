package bootstrap

import (
	"github.com/saegus/test-technique-romain-chenard/pkg/configu"
	"github.com/saegus/test-technique-romain-chenard/pkg/database"

	"github.com/saegus/test-technique-romain-chenard/pkg/routing"
)

func Serve() {
	configu.Set()

	database.Connect()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
