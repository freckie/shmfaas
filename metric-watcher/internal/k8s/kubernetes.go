package k8s

import (
	"context"

	"k8s.io/client-go/rest"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

type K8sClient struct {
	Ctx       context.Context
	MetricsCS *metricsv.Clientset
}

func LoadK8s() (K8sClient, error) {
	var client K8sClient

	config, err := rest.InClusterConfig()
	if err != nil {
		return client, err
	}

	ctx := context.TODO()
	metricsCS, err := metricsv.NewForConfig(config)
	if err != nil {
		return client, err
	}
	ctx = context.WithValue(ctx, "k8s_metrics_clientset", metricsCS)

	client = K8sClient{
		Ctx:       ctx,
		MetricsCS: metricsCS,
	}

	return client, nil
}
