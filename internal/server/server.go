package server

import (
	"scaler/internal/config"
	"scaler/internal/routes"
	"scaler/internal/services"
)

func SetupAndListen() error {
	cfg := config.GetServerConfigs()

	svcs, err := services.Factory(cfg)
	if err != nil {
		return err
	}

	return routes.GetApp(svcs, cfg).Listen(cfg.ServerAddress)
}
