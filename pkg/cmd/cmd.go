package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomschinelli/kummy/pkg/cmd/overview"
)

func NewDefaultKummyCommand() *cobra.Command {
	return NewkummyCommand()
}

func NewkummyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kummy",
		Short: "kummy - cluster for dummies",
		Long:  `kummy a tool targeting developer without or little cluster knowledge so that they can inspect their applications.`,
	}

	overviewCmd := overview.NewCmdOverview()

	cmd.AddCommand(overviewCmd)
	return cmd
}
