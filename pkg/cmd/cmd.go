package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tomschinelli/kummy/pkg/cmd/get"
	"github.com/tomschinelli/kummy/pkg/cmd/logs"
	"github.com/tomschinelli/kummy/pkg/cmd/overview"
)

var (
	Namespace string
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

	cmd.AddCommand(overview.NewCmdOverview())
	cmd.AddCommand(get.NewCmdGet())
	cmd.AddCommand(logs.NewCmdLogs())

	cmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "default", "Target Namespace")

	return cmd
}
