package migration

import (
	"fmt"
	"log"

	"github.com/saegus/test-technique-romain-chenard/data/database"
	Models "github.com/saegus/test-technique-romain-chenard/data/models"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&Models.User{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("migration done ...")
}
