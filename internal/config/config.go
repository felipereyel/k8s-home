package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	ServerAddress     string
	UseServiceAccount bool
	BasicPassword     string
}

func GetServerConfigs() ServerConfigs {
	config := ServerConfigs{}

	config.BasicPassword = os.Getenv("BASIC_PASSWORD")
	config.UseServiceAccount = (os.Getenv("USE_SA") == "true")

	envPort := os.Getenv("PORT")
	if envPort != "" {
		config.ServerAddress = ":" + envPort
	} else {
		config.ServerAddress = ":3000"
	}

	return config
}
