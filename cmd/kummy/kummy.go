package main

import (
	"github.com/tomschinelli/kummy/pkg/cmd"
	"os"
)

func main() {
	command := cmd.NewDefaultKummyCommand()
	if err := command.Execute(); err != nil {
		println(err)
		os.Exit(1)
	}
}
