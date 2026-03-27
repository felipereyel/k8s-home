package services

import (
	"k8s-home/internal/config"
	"k8s-home/internal/services/kubeclient"
)

type Services struct {
	KubeClient *kubeclient.KubeClient
}

func Factory(cfg config.ServerConfigs) (*Services, error) {
	kubeclientset, err := kubeclient.NewKubeClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Services{kubeclientset}, nil
}
