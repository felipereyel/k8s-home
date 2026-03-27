package routes

import (
	"scaler/internal/components"
	"scaler/internal/services"

	"github.com/gofiber/fiber/v2"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
)

func deploymentsDetails(svcs *services.Services, c *fiber.Ctx) error {
	c.Set("HX-Refresh", "true")
	namespace := c.Params("namespace")
	name := c.Params("deployment")

	d, err := svcs.KubeClient.GetDeployment(namespace, name)
	if err != nil {
		return err
	}

	svc, err := findServiceForApp(svcs, namespace, name)
	if err != nil {
		svc = nil
	}

	var ingresses []networkingv1.Ingress
	if svc != nil {
		ingresses = svcs.KubeClient.GetIngressForService(namespace, svc.Name)
	}

	return sendPage(c, components.DeploymentDetailsPage(d, svc, ingresses))
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

func findServiceForApp(svcs *services.Services, namespace, appName string) (*corev1.Service, error) {
	services, err := svcs.KubeClient.ListServices()
	if err != nil {
		return nil, err
	}

	for _, svc := range services {
		if svc.Namespace != namespace {
			continue
		}
		if svc.Spec.Selector != nil && svc.Spec.Selector["app"] == appName {
			return &svc, nil
		}
	}
	return nil, nil
}

func findStatefulSetServiceForApp(svcs *services.Services, namespace, appName string) (*corev1.Service, error) {
	services, err := svcs.KubeClient.ListServices()
	if err != nil {
		return nil, err
	}

	for _, svc := range services {
		if svc.Namespace != namespace {
			continue
		}
		if svc.Spec.Selector != nil && svc.Spec.Selector["app"] == appName {
			return &svc, nil
		}
	}
	return nil, nil
}
