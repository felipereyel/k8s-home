package routes

import (
	"scaler/internal/config"
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func GetApp(svcs *services.Services, cfg config.ServerConfigs) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	if cfg.BasicPassword != "" {
		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				"admin": cfg.BasicPassword,
			},
		}))
	}

	app.Get("/", bind(svcs, home))

	app.Get("/_deployments/:namespace/:deployment", bind(svcs, deploymentsDetails))
	app.Post("/_deployments/:namespace/:deployment", bind(svcs, deploymentsToggle))

	app.Get("/_statefulsets/:namespace/:statefulset", bind(svcs, statefulsetsDetails))
	app.Post("/_statefulsets/:namespace/:statefulset", bind(svcs, statefulsetsToggle))

	app.Use("/_statics", staticsHandler)
	app.Use("/_healthz", healthzHandler)
	app.Use(notFoundHandler)

	return app
}
