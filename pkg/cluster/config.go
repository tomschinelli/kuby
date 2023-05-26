package cluster

import (
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
)

// NewAutoConfig Auto find cluster config
//
// Resolves config in the following order:
//   - 1. configured via kummy
//   - 2. auto find with controller-runtime.GetConfig
func NewAutoConfig() (*rest.Config, error) {
	// todo configure config
	return ctrl.GetConfig()
}
