package logs

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tomschinelli/kummy/pkg/cluster"
	v1 "k8s.io/api/apps/v1"
	"strings"
)

func validDeployments() ([]v1.Deployment, error) {
	ctx := context.Background()
	clst := cluster.NewDefault()
	return clst.Deployments(ctx, "")
}

func validLogsArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	deployments, err := validDeployments()
	if err != nil {
		println(err)
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	var result []string
	for _, depl := range deployments {
		if toComplete == "" || strings.HasPrefix(depl.Name, toComplete) {
			result = append(result, depl.Name)
		}
	}
	return result, cobra.ShellCompDirectiveNoFileComp
}
