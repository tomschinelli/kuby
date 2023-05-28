package get

import "github.com/spf13/cobra"

func NewCmdGetLogs() *cobra.Command {

	return &cobra.Command{
		Use: "logs",
		Run: func(cmd *cobra.Command, args []string) {
			println("get logs - not implemented.")
		},
	}
}
