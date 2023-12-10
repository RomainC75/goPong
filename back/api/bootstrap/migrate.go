package bootstrap

import (
	"github.com/saegus/test-technique-romain-chenard/config"
	"github.com/saegus/test-technique-romain-chenard/data/database"
	"github.com/saegus/test-technique-romain-chenard/data/migration"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
