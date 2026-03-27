package routes

import (
	"scaler/internal/components"
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
)

func statefulsetsDetails(svcs *services.Services, c *fiber.Ctx) error {
	c.Set("HX-Refresh", "true")
	namespace := c.Params("namespace")
	name := c.Params("statefulset")

	s, err := svcs.KubeClient.GetStatefulSet(namespace, name)
	if err != nil {
		return err
	}

	return sendPage(c, components.StatefulsetDetailsPage(s))
}

func statefulsetsToggle(svcs *services.Services, c *fiber.Ctx) error {
	c.Set("HX-Refresh", "true")
	namespace := c.Params("namespace")
	name := c.Params("statefulset")

	s, err := svcs.KubeClient.GetStatefulSet(namespace, name)
	if err != nil {
		return err
	}

	replicas := 0
	if s.Spec.Replicas != nil && *s.Spec.Replicas == 0 {
		replicas = 1
	}

	if err = svcs.KubeClient.ScaleStatefulSet(s, int32(replicas)); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
