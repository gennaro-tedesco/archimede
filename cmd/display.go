package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func displayInfo(fileFormat string) {
	filesList := getFiles()
	totalFiles, extCount := countFiles(filesList)
	total := totalFiles["normal"] + totalFiles["hidden"]

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.Style().Color.Row = text.Colors{text.FgCyan}
	t.Style().Options.SeparateColumns = false
	t.Style().Box.BottomLeft = "╰"
	t.Style().Box.TopLeft = "╭"
	t.Style().Box.TopRight = "╮"
	t.Style().Box.BottomRight = "╯"

	if fileFormat == "long" {
		t.AppendRow(table.Row{
			"Files:", fmt.Sprintf("%v regular + %v hidden (%v%% %v, %v%% %v, %v%% %v)",
				totalFiles["normal"], totalFiles["hidden"],
				100*extCount[0].Value/total, extCount[0].Key,
				100*extCount[1].Value/total, extCount[1].Key,
				100*extCount[2].Value/total, extCount[2].Key),
		})
	} else {
		t.AppendRow(table.Row{
			"Files:", fmt.Sprintf("%v + %v ", totalFiles["normal"], totalFiles["hidden"]),
		})
	}
	t.AppendSeparator()
	t.Render()
}
