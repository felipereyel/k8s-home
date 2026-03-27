package routes

import (
	"scaler/internal/components"
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
)

func home(svcs *services.Services, c *fiber.Ctx) error {
	deployments, err := svcs.KubeClient.ListDeployments()
	if err != nil {
		return err
	}

	statefulsets, err := svcs.KubeClient.ListStatefulSets()
	if err != nil {
		return err
	}

	return sendPage(c, components.AppListPage(deployments, statefulsets))
}
