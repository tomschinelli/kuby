package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Cluster struct {
	clientset *kubernetes.Clientset
}

func New(config *rest.Config) (*Cluster, error) {

	clientset, err := kubernetes.NewForConfig(config)
	cluster := &Cluster{
		clientset: clientset,
	}
	return cluster, err
}

func NewDefault() (*Cluster, error) {
	config, err := NewAutoConfig()
	if err != nil {
		return nil, err
	}
	return New(config)
}
