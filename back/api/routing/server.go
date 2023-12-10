package routing

import (
	"fmt"
	"log"

	"github.com/saegus/test-technique-romain-chenard/config"
)

func Serve() {
	configs := config.Get()

	r := GetRouter()

	err := r.Run(fmt.Sprintf(":%v", configs.Server.Port))

	if err != nil {
		log.Fatal("Error in routing !")
		return
	}
}
