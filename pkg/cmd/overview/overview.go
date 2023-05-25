package overview

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

func NewCmdOverview() *cobra.Command {
	return &cobra.Command{
		Use: "overview",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			config, err := ctrl.GetConfig()
			if err != nil {
				return err
			}
			clientset := kubernetes.NewForConfigOrDie(config)
			namespace := "" // all namespaces
			items, err := GetPods(clientset, ctx, namespace)
			if err != nil {
				return err
			}
			for _, item := range items {
				fmt.Printf("%+v\n", item.Name)
			}
			return nil
		},
	}
}

func GetPods(clientset *kubernetes.Clientset, ctx context.Context, namespace string) ([]v1.Pod, error) {

	list, err := clientset.CoreV1().Pods(namespace).
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}
