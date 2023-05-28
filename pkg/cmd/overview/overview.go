package overview

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tomschinelli/kummy/pkg/cluster"
	"github.com/tomschinelli/kummy/pkg/overview"
	v1 "k8s.io/api/core/v1"
)

func NewCmdOverview() *cobra.Command {
	return &cobra.Command{
		Use: "overview",
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := getData(cmd.Flags())
			if err != nil {
				return err
			}
			PrintTable(data)
			return nil
		},
	}
}
func getData(flags *pflag.FlagSet) ([]v1.Pod, error) {
	namespace, _ := flags.GetString("namespace")

	ctx := context.Background()
	clst := cluster.NewDefault()
	options := overview.NewDefaultOptions()
	options.Namespace = namespace
	o := overview.New(clst, options)
	return o.Get(ctx)

}
