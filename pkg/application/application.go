package application

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"golang.org/x/exp/slices"
)

import v1 "k8s.io/api/apps/v1"

func PrintForConfig(config *Config) error {

	applications, err := listSimplified(config)
	if err != nil {
		return err
	}
	switch config.outputFormat {
	case "table":
		printTable(applications)
	}

	return nil
}

func listSimplified(config *Config) ([]v1.Deployment, error) {
	if config.err != nil {
		return nil, config.err
	}
	deployments, err := config.cluster.Deployments(config.ctx, "")
	if err != nil {
		return nil, err
	}
	hiddenNamespaces := []string{
		"kube-system",
	}
	var result []v1.Deployment
	for _, depl := range deployments {

		if slices.Contains(hiddenNamespaces, depl.Namespace) {
			continue
		}
		result = append(result, depl)
	}

	return result, nil
}

func printTable(deployments []v1.Deployment) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Name")
	tbl.
		WithHeaderFormatter(headerFmt).
		WithFirstColumnFormatter(columnFmt)

	for _, item := range deployments {
		tbl.AddRow(
			item.Name,
		)
	}
	tbl.Print()
}
