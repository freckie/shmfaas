package k8s

import (
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	klog "k8s.io/klog/v2"

	iconfig "github.com/freckie/shmfaas/event-watcher/internal/config"
)

// WatchEvents
//  inspired from https://stackoverflow.com/a/49231503
func (c *K8sClient) WatchEvents(stopSig chan struct{}, cfg *iconfig.WatcherConfig) {
	cs := c.clientset

	filter := filterFunc(cfg)

	targets := cache.NewListWatchFromClient(
		cs.EventsV1().RESTClient(),
		"events",
		cfg.Namespace,
		fields.Everything(),
	)
	_, controller := cache.NewInformer(
		targets,
		&eventsv1.Event{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				event := filter(obj.(*eventsv1.Event))
				if event != nil {
					ts := event.EventTime.Format(metav1.RFC3339Micro)
					ts2 := event.CreationTimestamp.Format(metav1.RFC3339Micro)
					klog.Infof("[Added] %s/%s\n > EventTime: %s (%s)\n > Reason: %s\n\n",
						event.Regarding.Kind,
						event.Regarding.Name,
						ts,
						ts2,
						event.Reason,
					)
				}
			},
			DeleteFunc: func(obj interface{}) {
				klog.Infof("[Deleted]\n")
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				klog.Infof("[Updated]\n")
			},
		},
	)

	controller.Run(stopSig)
}

func filterFunc(cfg *iconfig.WatcherConfig) func(*eventsv1.Event) *eventsv1.Event {
	targets := cfg.Targets
	flag := cfg.FilterTargets

	return func(event *eventsv1.Event) *eventsv1.Event {
		if flag {
			for _, it := range targets {
				if event.Reason == it.Reason && event.Regarding.Kind == it.ResourceKind {
					return event
				}
			}
			return nil
		}
		return event
	}
}
