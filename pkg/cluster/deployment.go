package cluster

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Cluster) GetDeployments(ctx context.Context, namespace string) ([]v1.Deployment, error) {
	if c.err != nil {
		return nil, c.err
	}
	list, err := c.clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list.Items, nil

}
