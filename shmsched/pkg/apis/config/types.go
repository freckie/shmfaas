package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ShmScoringArgs configures ShmScoring plugin
type ShmScoringArgs struct {
	metav1.TypeMeta

	// List of {addr:port}s of the shmm daemonsets
	AddrPorts []string
}
