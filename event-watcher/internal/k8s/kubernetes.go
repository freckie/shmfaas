package k8s

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sClient struct {
	ctx       context.Context
	clientset *kubernetes.Clientset
}

func LoadK8s() (K8sClient, error) {
	var client K8sClient

	config, err := rest.InClusterConfig()
	if err != nil {
		return client, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return client, err
	}

	client = K8sClient{
		ctx:       context.Background(),
		clientset: clientset,
	}

	return client, nil
}
