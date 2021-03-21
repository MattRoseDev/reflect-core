package setting

import (
	"log"

	"github.com/favecode/reflect-core/config"
	"github.com/spf13/viper"
)


func Get() (*config.Configuration) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &configuration
}
