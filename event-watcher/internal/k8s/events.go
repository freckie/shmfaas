package k8s

import (
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	klog "k8s.io/klog/v2"
)

// WatchEvents
//  inspired from https://stackoverflow.com/a/49231503
func (c *K8sClient) WatchEvents(stopSig chan struct{}, namespace string) {
	cs := c.clientset

	targets := cache.NewListWatchFromClient(
		cs.EventsV1().RESTClient(),
		"events",
		namespace,
		fields.Everything(),
	)
	_, controller := cache.NewInformer(
		targets,
		&eventsv1.Event{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				event := obj.(*eventsv1.Event)
				ts := event.EventTime.Format(metav1.RFC3339Micro)
				ts2 := event.CreationTimestamp.Format(metav1.RFC3339Micro)
				klog.Infof("[Added] %s/%s\n > EventTime: %s (%s)\n > Reason: %s\n > Event: %s\n\n",
					event.Regarding.Kind,
					event.Regarding.Name,
					ts,
					ts2,
					event.Reason,
					obj,
				)
			},
			DeleteFunc: func(obj interface{}) {
				event := obj.(*eventsv1.Event)
				ts := event.EventTime.Format(metav1.RFC3339Micro)
				ts2 := event.CreationTimestamp.Format(metav1.RFC3339Micro)
				klog.Infof("[Deleted] %s/%s\n > EventTime: %s (%s)\n > Reason: %s\n > Event: %s\n\n",
					event.Regarding.Kind,
					event.Regarding.Name,
					ts,
					ts2,
					event.Reason,
					obj,
				)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				klog.Infof("[Updated]\n")
			},
		},
	)

	controller.Run(stopSig)
}
