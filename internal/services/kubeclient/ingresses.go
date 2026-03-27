package kubeclient

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kubeclient *KubeClient) ListIngresses() ([]networkingv1.Ingress, error) {
	ingressList, err := kubeclient.clientset.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return ingressList.Items, nil
}

func (kubeclient *KubeClient) GetIngressForService(namespace, serviceName string) []networkingv1.Ingress {
	ingresses, err := kubeclient.ListIngresses()
	if err != nil {
		return nil
	}

	var matching []networkingv1.Ingress
	for _, ing := range ingresses {
		for _, rule := range ing.Spec.Rules {
			if rule.HTTP != nil {
				for _, path := range rule.HTTP.Paths {
					if path.Backend.Service != nil && path.Backend.Service.Name == serviceName {
						matching = append(matching, ing)
					}
				}
			}
		}
	}

	return matching
}
