package routes

import (
	"scaler/internal/components"
	"scaler/internal/services"
	"scaler/internal/utils"

	"github.com/gofiber/fiber/v2"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
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

	services, err := svcs.KubeClient.ListServices()
	if err != nil {
		return err
	}

	ingresses, err := svcs.KubeClient.ListIngresses()
	if err != nil {
		return err
	}

	hostname := c.Hostname()
	apps := buildApps(deployments, statefulsets, services, ingresses, hostname)

	return sendPage(c, components.AppListPage(apps))
}

func buildApps(deployments []v1.Deployment, statefulsets []v1.StatefulSet, services []corev1.Service, ingresses []networkingv1.Ingress, hostname string) []utils.App {
	serviceMap := make(map[string][]string)
	for _, svc := range services {
		if svc.Spec.Selector == nil {
			continue
		}
		key := svc.Namespace + "/" + svc.Spec.Selector["app"]
		serviceMap[key] = append(serviceMap[key], svc.Name)
	}

	ingressMap := make(map[string][]networkingv1.Ingress)
	for _, ing := range ingresses {
		for _, rule := range ing.Spec.Rules {
			if rule.HTTP != nil {
				for _, path := range rule.HTTP.Paths {
					if path.Backend.Service != nil {
						key := ing.Namespace + "/" + path.Backend.Service.Name
						ingressMap[key] = append(ingressMap[key], ing)
					}
				}
			}
		}
	}

	var apps []utils.App

	for _, d := range deployments {
		svcNames := serviceMap[d.Namespace+"/"+d.Name]
		var matchingIngs []networkingv1.Ingress
		for _, svcName := range svcNames {
			matchingIngs = append(matchingIngs, ingressMap[d.Namespace+"/"+svcName]...)
		}
		filteredIngs := utils.FilterIngressesByDomain(matchingIngs, hostname)
		app := utils.NewAppFromDeployment(d, filteredIngs)
		apps = append(apps, *app)
	}

	for _, s := range statefulsets {
		svcNames := serviceMap[s.Namespace+"/"+s.Name]
		var matchingIngs []networkingv1.Ingress
		for _, svcName := range svcNames {
			matchingIngs = append(matchingIngs, ingressMap[s.Namespace+"/"+svcName]...)
		}
		filteredIngs := utils.FilterIngressesByDomain(matchingIngs, hostname)
		app := utils.NewAppFromStatefulSet(s, filteredIngs)
		apps = append(apps, *app)
	}

	return apps
}
