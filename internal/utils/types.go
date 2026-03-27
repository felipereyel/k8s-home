package utils

import (
	"strconv"
	"strings"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
)

type AppType string

const (
	AppTypeDeployment  AppType = "deployment"
	AppTypeStatefulSet AppType = "statefulset"
)

type App struct {
	Name        string
	Namespace   string
	Type        AppType
	Replicas    int32
	Ingresses   []networkingv1.Ingress
	HostNetwork string
}

func NewAppFromDeployment(d v1.Deployment, ingresses []networkingv1.Ingress) *App {
	replicas := int32(0)
	if d.Spec.Replicas != nil {
		replicas = *d.Spec.Replicas
	}
	hostNetwork := getHostNetwork(d.Spec.Template.Spec)
	return &App{
		Name:        d.Name,
		Namespace:   d.Namespace,
		Type:        AppTypeDeployment,
		Replicas:    replicas,
		Ingresses:   ingresses,
		HostNetwork: hostNetwork,
	}
}

func NewAppFromStatefulSet(s v1.StatefulSet, ingresses []networkingv1.Ingress) *App {
	replicas := int32(0)
	if s.Spec.Replicas != nil {
		replicas = *s.Spec.Replicas
	}
	hostNetwork := getHostNetwork(s.Spec.Template.Spec)
	return &App{
		Name:        s.Name,
		Namespace:   s.Namespace,
		Type:        AppTypeStatefulSet,
		Replicas:    replicas,
		Ingresses:   ingresses,
		HostNetwork: hostNetwork,
	}
}

func getHostNetwork(spec corev1.PodSpec) string {
	if !spec.HostNetwork {
		return ""
	}
	for _, c := range spec.Containers {
		if len(c.Ports) > 0 {
			return strconv.Itoa(int(c.Ports[0].ContainerPort))
		}
	}
	return ""
}

func Int32Compare(a *int32, b int) bool {
	if a == nil {
		return false
	}
	return *a == int32(b)
}

func GetRootDomain(hostname string) string {
	parts := strings.Split(hostname, ".")
	if len(parts) >= 2 {
		return strings.Join(parts[len(parts)-2:], ".")
	}
	return ""
}

func FilterIngressesByDomain(ingresses []networkingv1.Ingress, hostname string) []networkingv1.Ingress {
	rootDomain := GetRootDomain(hostname)
	if rootDomain == "" {
		return ingresses
	}

	var filtered []networkingv1.Ingress
	for _, ing := range ingresses {
		for _, rule := range ing.Spec.Rules {
			if strings.HasSuffix(rule.Host, "."+rootDomain) || rule.Host == rootDomain {
				filtered = append(filtered, ing)
				break
			}
		}
	}
	return filtered
}
