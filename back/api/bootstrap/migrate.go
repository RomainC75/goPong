package bootstrap

import (
	"github.com/saegus/test-technique-romain-chenard/internal/database/migration"
	"github.com/saegus/test-technique-romain-chenard/pkg/configu"
	"github.com/saegus/test-technique-romain-chenard/pkg/database"
)

func Migrate() {
	configu.Set()

	database.Connect()

	migration.Migrate()
}
