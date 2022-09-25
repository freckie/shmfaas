package main

import (
	"os"

	ik8s "github.com/freckie/shmfaas/event-watcher/internal/k8s"
	klog "k8s.io/klog/v2"
)

func main() {
	// Environment variables
	namespace := os.Getenv("NAMESPACE")
	klog.Infof("Target namespace : %s\n", namespace)

	client, err := ik8s.LoadK8s()
	if err != nil {
		klog.Errorf("Failed to load kubernetes client. %s", err.Error())
	}

	stopSig := make(chan struct{})
	client.WatchEvents(stopSig, namespace)
}
