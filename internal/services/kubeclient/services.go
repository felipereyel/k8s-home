package kubeclient

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kubeclient *KubeClient) ListServices() ([]corev1.Service, error) {
	serviceList, err := kubeclient.clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return serviceList.Items, nil
}
