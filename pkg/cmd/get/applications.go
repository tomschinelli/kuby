package get

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tomschinelli/kummy/pkg/application"
	"github.com/tomschinelli/kummy/pkg/cluster"
)

func NewCmdGetApplications() *cobra.Command {
	return &cobra.Command{
		Use: "applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			clst := cluster.NewDefault()
			config := application.
				ParseConfigForCommand(cmd).
				UseContext(ctx).
				UseCluster(clst)
			err := application.PrintForConfig(config)
			return err
		},
	}
}
