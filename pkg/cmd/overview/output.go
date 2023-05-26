package overview

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	v1 "k8s.io/api/core/v1"
)

func PrintTable(rows []v1.Pod) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Name", "Status")
	tbl.
		WithHeaderFormatter(headerFmt).
		WithFirstColumnFormatter(columnFmt)

	for _, item := range rows {
		tbl.AddRow(
			item.Name,
			item.Status.Phase)
	}
	tbl.Print()

}
