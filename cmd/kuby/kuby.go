package main

import (
	"github.com/tomschinelli/kuby/pkg/cmd"
	"os"
)

func main() {
	command := cmd.NewDefaultKubyCommand()
	if err := command.Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}
