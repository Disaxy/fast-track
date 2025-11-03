package main

import (
	"github.com/Disaxy/fast-track/tree/main/env-variables/config"
	"github.com/Disaxy/fast-track/tree/main/env-variables/server"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
)

func main() {
	var cfg config.Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		slog.Error(err.Error())
	}

	server.Serve(&cfg)
}
