package main

import (
	"github.com/VashUber/coursework/golang/internal/app"
	"github.com/VashUber/coursework/golang/internal/config"
)

func main() {
	cfg := config.NewConfig()

	app.Run(cfg)
}
