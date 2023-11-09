package migration

import (
	"fmt"
	"log"

	listModels "github.com/saegus/test-technique-romain-chenard/internal/modules/list/models"
	taskModels "github.com/saegus/test-technique-romain-chenard/internal/modules/task/models"
	userModels "github.com/saegus/test-technique-romain-chenard/internal/modules/user/models"
	"github.com/saegus/test-technique-romain-chenard/pkg/database"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &listModels.List{}, &taskModels.Task{})
	

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("migration done ...")
}
