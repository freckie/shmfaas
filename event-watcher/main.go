package main

import (
	"os"

	iconfig "github.com/freckie/shmfaas/event-watcher/internal/config"
	ik8s "github.com/freckie/shmfaas/event-watcher/internal/k8s"
	klog "k8s.io/klog/v2"
)

func main() {
	// config
	cfgFilename := os.Getenv("CONFIG")
	cfg, err := iconfig.LoadConfig(cfgFilename)
	if err != nil {
		klog.Errorf("Failed to load event-watcher-config. %s", err.Error())
		panic("failed to load config")
	}
	if cfg.FilterTargets {
		klog.Infoln("[ Filtering Targets ]")
		for _, it := range cfg.Targets {
			klog.Infof(": %s (%s)", it.Reason, it.ResourceKind)
		}
	}

	client, err := ik8s.LoadK8s()
	if err != nil {
		klog.Errorf("Failed to load kubernetes client. %s", err.Error())
		panic("failed to load kubernetes client")
	}

	stopSig := make(chan struct{})
	client.WatchEvents(stopSig, cfg)
}
