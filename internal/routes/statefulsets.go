package routes

import (
	"k8s-home/internal/components"
	"k8s-home/internal/services"

	"github.com/gofiber/fiber/v2"
	networkingv1 "k8s.io/api/networking/v1"
)

func statefulsetsDetails(svcs *services.Services, c *fiber.Ctx) error {
	c.Set("HX-Refresh", "true")
	namespace := c.Params("namespace")
	name := c.Params("statefulset")

	s, err := svcs.KubeClient.GetStatefulSet(namespace, name)
	if err != nil {
		return err
	}

	svc, err := findStatefulSetServiceForApp(svcs, namespace, name)
	if err != nil {
		svc = nil
	}

	var ingresses []networkingv1.Ingress
	if svc != nil {
		ingresses = svcs.KubeClient.GetIngressForService(namespace, svc.Name)
	}

	var hostPorts []string
	if s.Spec.Template.Spec.HostNetwork {
		hostPorts = extractContainerPorts(s.Spec.Template.Spec.Containers)
	}

	return sendPage(c, components.StatefulsetDetailsPage(s, hostPorts, ingresses))
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
