package kubeclient

import (
	"path/filepath"
	"scaler/internal/config"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubeClient struct {
	clientset *kubernetes.Clientset
}

func NewKubeClient(cfg config.ServerConfigs) (*KubeClient, error) {
	var config *rest.Config
	var err error

	if cfg.UseServiceAccount {
		config, err = rest.InClusterConfig()
	} else {
		kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	}

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &KubeClient{clientset}, nil
}
