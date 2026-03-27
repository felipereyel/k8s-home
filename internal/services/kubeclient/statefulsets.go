package kubeclient

import (
	"context"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kubeclient *KubeClient) GetStatefulSet(namespace, name string) (*v1.StatefulSet, error) {
	return kubeclient.clientset.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (kubeclient *KubeClient) ListStatefulSets() ([]v1.StatefulSet, error) {
	statefulSetList, err := kubeclient.clientset.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	statefulSets := make([]v1.StatefulSet, 0)
	for _, statefulSet := range statefulSetList.Items {
		annotations := statefulSet.GetAnnotations()
		if annotations == nil {
			continue
		}

		if annotations["scaler.reyel.cloud/enabled"] == "true" {
			statefulSets = append(statefulSets, statefulSet)
		}
	}

	return statefulSets, nil
}

func (kubeclient *KubeClient) ScaleStatefulSet(statefulSet *v1.StatefulSet, replicas int32) error {
	statefulSet.Spec.Replicas = &replicas

	_, err := kubeclient.clientset.AppsV1().StatefulSets(statefulSet.Namespace).Update(context.TODO(), statefulSet, metav1.UpdateOptions{})

	return err
}
