package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load env")
	}

	cfg := &Config{
		Http: HttpConfig{
			Port: os.Getenv("PORT"),
		},

		Controller: ControllerConfig{
			StaticDir:   "./static",
			SpaViewRoot: os.Getenv("SPA_ROOT"),
		},
	}

	return cfg
}
