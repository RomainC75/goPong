package database

import (
	"fmt"
	"log"

	"github.com/saegus/test-technique-romain-chenard/pkg/configu"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	cfg := configu.Get()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
		return
	}

	DB = db
}
