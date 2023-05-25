package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomschinelli/kuby/pkg/cmd/overview"
)

func NewDefaultKubyCommand() *cobra.Command {
	return NewKubyCommand()
}

func NewKubyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kuby",
		Short: "Kuby - kubernetes for dummies",
		Long:  `Kuby a tool targeting developer without or little kubernetes knowledge so that they can inspect their applications.`,
	}

	overviewCmd := overview.NewCmdOverview()

	cmd.AddCommand(overviewCmd)
	return cmd
}
