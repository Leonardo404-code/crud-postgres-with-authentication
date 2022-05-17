package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadDotEnv load all enviroment variables
func LoadDotEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env.local")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading env file, %s", err)
	}
}

// LoadDotEnvTest is a especial function for load all enviroment variables in tests
func LoadDotEnvTests() {
	viper.AddConfigPath("../../")
	viper.SetConfigName(".env.local")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading env file, %s", err)
	}
}
