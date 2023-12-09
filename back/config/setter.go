package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading config")
	}

	err := viper.Unmarshal(&configurations)
	fmt.Print("BEFOS : ", configurations)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}

}
