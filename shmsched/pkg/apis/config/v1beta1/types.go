// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:defaulter-gen=true

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ShmScoringArgs struct {
	metav1.TypeMeta `json:",inline"`

	// List of {addr:port}s of the shmm daemonsets
	AddrPorts *[]string `json:"addrports,omitempty"`
}
