package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

var flagEnvPath string

func ParseFlags() {
	flag.StringVar(&flagEnvPath, "env-path", ".", "path to .env file")
	flag.Parse()
}

// LoadDotEnv load all enviroment variables
func LoadDotEnv() {
	viper.AddConfigPath(flagEnvPath)
	viper.SetConfigName(".env.local")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading env file, %s", err)
	}
}

func LoadDotEnvTests() {
	viper.AddConfigPath("../../")
	viper.SetConfigName(".env.local")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading env file, %s", err)
	}
}
