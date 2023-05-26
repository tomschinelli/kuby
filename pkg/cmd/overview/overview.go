package overview

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/tomschinelli/kummy/pkg/cluster"
	"github.com/tomschinelli/kummy/pkg/overview"
	v1 "k8s.io/api/core/v1"
)

func NewCmdOverview() *cobra.Command {
	return &cobra.Command{
		Use: "overview",
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := getData()
			if err != nil {
				return err
			}
			PrintTable(data)
			return nil
		},
	}
}
func getData() ([]v1.Pod, error) {
	ctx := context.Background()
	clst, err := cluster.NewDefault()
	if err != nil {
		return nil, err
	}

	options := overview.NewDefaultOptions()
	options.AllNamespace = true
	o := overview.New(clst, options)

	return o.Get(ctx)

}
