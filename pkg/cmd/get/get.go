package get

import (
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {

	cmd := &cobra.Command{
		Use: "get",
	}

	cmd.AddCommand(NewCmdGetApplications())
	return cmd
}
