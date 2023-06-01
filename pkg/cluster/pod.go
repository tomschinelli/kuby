package cluster

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (c *Cluster) GetPods(ctx context.Context, namespace string) ([]v1.Pod, error) {
	if c.err != nil {
		return nil, c.err
	}
	list, err := c.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}

func (c *Cluster) Pods(ctx context.Context, namespace string, labels map[string]string) (*v1.PodList, error) {
	if c.err != nil {
		return nil, c.err
	}
	var selectors []string
	for key, val := range labels {
		selectors = append(selectors, fmt.Sprintf("%s=%s", key, val))
	}
	labelSelector := strings.Join(selectors, ",")
	return c.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
}
