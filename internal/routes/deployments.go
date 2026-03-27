package routes

import (
	"scaler/internal/components"
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
)

func deploymentsDetails(svcs *services.Services, c *fiber.Ctx) error {
	c.Set("HX-Refresh", "true")
	namespace := c.Params("namespace")
	name := c.Params("deployment")

	d, err := svcs.KubeClient.GetDeployment(namespace, name)
	if err != nil {
		return err
	}

	return sendPage(c, components.DeploymentDetailsPage(d))
}

func deploymentsToggle(svcs *services.Services, c *fiber.Ctx) error {
	c.Set("HX-Refresh", "true")
	namespace := c.Params("namespace")
	name := c.Params("deployment")

	d, err := svcs.KubeClient.GetDeployment(namespace, name)
	if err != nil {
		return err
	}

	replicas := 0
	if d.Spec.Replicas != nil && *d.Spec.Replicas == 0 {
		replicas = 1
	}

	if err = svcs.KubeClient.ScaleDeployment(d, int32(replicas)); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
