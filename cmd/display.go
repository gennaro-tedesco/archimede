package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/term"
)

func colourMap() map[string]text.Color {
	colourMap := map[string]text.Color{
		"black":   text.FgBlack,
		"cyan":    text.FgCyan,
		"green":   text.FgGreen,
		"yellow":  text.FgYellow,
		"blue":    text.FgBlue,
		"magenta": text.FgMagenta,
		"red":     text.FgRed,
		"white":   text.FgWhite,
	}
	return colourMap
}

func displayInfo(
	fileFormat string,
	textColour string,
	separator string,
	gitFolder bool) {

	t := initTable(textColour)

	filesList := getFiles(gitFolder)
	totalFiles, extCount := countFiles(filesList)
	total := totalFiles["normal"] + totalFiles["hidden"]
	if fileFormat == "long" && len(extCount) > 2 {
		t.AppendRow(table.Row{
			fmt.Sprintf("Files%v", separator), fmt.Sprintf("%v regular + %v hidden (%v%% %v, %v%% %v, %v%% %v)",
				totalFiles["normal"], totalFiles["hidden"],
				100*extCount[0].Value/total, extCount[0].Key,
				100*extCount[1].Value/total, extCount[1].Key,
				100*extCount[2].Value/total, extCount[2].Key),
		})
	} else {
		t.AppendRow(table.Row{
			fmt.Sprintf("Files%v", separator), fmt.Sprintf("%v + %v ", totalFiles["normal"], totalFiles["hidden"]),
		})
	}

	dirsList := countDirs()
	if fileFormat == "long" {
		t.AppendRow(table.Row{
			fmt.Sprintf("Dirs%v", separator), fmt.Sprintf("%v   + %v  /  + %v  / /  ",
				dirsList["one"], dirsList["two"], dirsList["three"]),
		})
	} else {
		t.AppendRow(table.Row{
			fmt.Sprintf("Dirs%v", separator), fmt.Sprintf("%v + %v +%v ", dirsList["one"], dirsList["two"], dirsList["three"]),
		})
	}

	t.AppendSeparator()
	t.Render()
}

func initTable(textColour string) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	width, _, _ := term.GetSize(0)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, WidthMax: width / 3},
		{Number: 2, WidthMax: 2 * width / 3},
	})

	if colour, ok := colourMap()[textColour]; ok {
		t.Style().Color.Row = text.Colors{colour}
	} else {
		t.Style().Color.Row = text.Colors{text.FgWhite}
	}
	t.Style().Options.SeparateColumns = false
	t.Style().Box.BottomLeft = "╰"
	t.Style().Box.TopLeft = "╭"
	t.Style().Box.TopRight = "╮"
	t.Style().Box.BottomRight = "╯"
	return t
}
