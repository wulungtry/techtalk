package persist

import (
	"github.com/spf13/viper"
	"log"
)

//ViperInit initialize viper configuration
func ViperInit() {
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No configuration file found")
	}
}
