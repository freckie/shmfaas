// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ShmScoringArgs struct {
	metav1.TypeMeta

	// List of {addr:port}s of the shmm daemonsets
	AddrPorts []string
}
