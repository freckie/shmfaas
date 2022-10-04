package k8s

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func GetPodsMemoryUsage(ctx context.Context, namespace string) (int64, error) {
	cs := ctx.Value("k8s_metrics_clientset")

	metrices, err := cs.(*metricsv.Clientset).MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return 0, err
	}

	var mem int64 = 0
	for _, m := range metrices.Items {
		for _, c := range m.Containers {
			mem += c.Usage.Memory().Value()
		}
	}

	return mem, nil
}
