package logs

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tomschinelli/kummy/pkg/cluster"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

var (
	ctx         context.Context
	clst        *cluster.Cluster
	logsExample = `
	# Show logs for application
	kummy logs <application>

	# Select application and show logs
	kummy logs
`
)

// NewCmdLogs deliver the log output for all pods of a deployment
//
// TODO WIP
//
//	todo get out of prove of concept state (structure and polish)
//	todo ask namespace for multiple deployment occurrences
//
// Usage:
//
//	> kummy logs deployment_name
//	> kummy logs
//
// If no deployment specified, it asks you to select one from a list.
// If deployment is specified, but it exists in multiple namespaces, it asks
// to select it.
func NewCmdLogs() *cobra.Command {

	return &cobra.Command{
		Use:               "logs",
		Args:              cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
		ValidArgsFunction: validLogsArgs,
		Example:           logsExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			initCluster()
			deployment, err := resolveDeployment(args)
			pods, err := resolvePods(deployment)
			println(pods)

			return err
		},
	}
}

func initCluster() {
	ctx = context.Background()
	clst = cluster.NewDefault()
}
func resolveDeployment(args []string) (*appsv1.Deployment, error) {
	var deploymentName string
	var namespace string
	var err error
	if len(args) == 1 {
		deploymentName = args[0]
		namespace = ""
	} else {
		deploymentName, namespace, err = promptDeployment()
	}
	deployment, err := clst.Deployment(ctx, namespace, deploymentName)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

func resolvePods(deployment *appsv1.Deployment) (*corev1.PodList, error) {
	labels := deployment.Spec.Template.Labels
	return clst.Pods(ctx, deployment.Namespace, labels)
}
