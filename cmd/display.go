package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func displayInfo() {
	filesList := getFiles()
	totalFiles, _ := countFiles(filesList)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRows([]table.Row{
		{"Files:", fmt.Sprintf("%v regular + %v hidden", totalFiles["normal"], totalFiles["hidden"])},
	})
	t.AppendSeparator()
	t.SetStyle(table.StyleLight)
	t.Style().Color.Row = text.Colors{text.FgCyan}
	t.Style().Options.SeparateColumns = false
	t.Render()
}
