package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:defaulter-gen=true

// ShmScoringArgs configures ShmScoring plugin
type ShmScoringArgs struct {
	metav1.TypeMeta `json:",inline"`

	// List of {addr:port}s of the shmm daemonsets
	AddrPorts *[]string `json:"addrports,omitempty"`
}
