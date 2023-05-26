package overview

import (
	"context"
	"github.com/tomschinelli/kummy/pkg/cluster"
	v1 "k8s.io/api/core/v1"
)

type Overview struct {
	cluster *cluster.Cluster
	Options *Options
}

func (o *Overview) Get(ctx context.Context) ([]v1.Pod, error) {
	return o.cluster.GetPods(ctx, o.GetNamespace())

}

func (o *Overview) GetNamespace() string {
	if o.Options.AllNamespace {
		return ""
	}
	return o.Options.Namespace
}

func New(cluster *cluster.Cluster, options *Options) *Overview {
	return &Overview{
		cluster: cluster,
		Options: options,
	}
}
