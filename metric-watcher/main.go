package main

import (
	"fmt"
	"net/http"
	"os"

	iconfig "github.com/freckie/shmfaas/metrics-watcher/internal/config"
	ik8s "github.com/freckie/shmfaas/metrics-watcher/internal/k8s"
	klog "k8s.io/klog/v2"
)

type MetricsHandler struct {
	client *ik8s.K8sClient
	cfg    *iconfig.WatcherConfig
}

func (mh *MetricsHandler) handler(w http.ResponseWriter, r *http.Request) {
	val, err := ik8s.GetPodsMemoryUsage(mh.client.Ctx, mh.cfg.Namespace)
	if err != nil {
		klog.Error(err)
		return
	}

	result := fmt.Sprintf("x_pod_mem_usage{namespace=\"%s\"} %d", mh.cfg.Namespace, val)
	fmt.Fprintf(w, result)
}

func main() {
	// config
	cfgFilename := os.Getenv("CONFIG")
	cfg, err := iconfig.LoadConfig(cfgFilename)
	if err != nil {
		klog.Errorf("Failed to load metric-watcher-config. %s", err.Error())
		panic("failed to load config")
	}

	client, err := ik8s.LoadK8s()
	if err != nil {
		klog.Errorf("Failed to load kubernetes client. %s", err.Error())
		panic("failed to load kubernetes client")
	}

	handler := MetricsHandler{
		client: &client,
		cfg:    cfg,
	}
	http.HandleFunc("/metrics", handler.handler)
	klog.Info("Listening on 8000 ..")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		klog.Fatal("ListenAndServe:", err)
	}
}
