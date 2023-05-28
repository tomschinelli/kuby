package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Cluster struct {
	err       error
	clientset *kubernetes.Clientset
}

func New(config *rest.Config) *Cluster {
	cluster := &Cluster{}

	clientset, err := kubernetes.NewForConfig(config)
	cluster.clientset = clientset
	cluster.err = err

	return cluster
}

func NewDefault() *Cluster {
	config, err := NewAutoConfig()
	if err != nil {
		return &Cluster{err: err}
	}
	return New(config)
}
