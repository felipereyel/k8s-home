package server

import (
	"k8s-home/internal/config"
	"k8s-home/internal/routes"
	"k8s-home/internal/services"
)

func SetupAndListen() error {
	cfg := config.GetServerConfigs()

	svcs, err := services.Factory(cfg)
	if err != nil {
		return err
	}

	return routes.GetApp(svcs, cfg).Listen(cfg.ServerAddress)
}
