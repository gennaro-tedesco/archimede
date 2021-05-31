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
	printFormat string,
	textColour string,
	separator string,
	gitFolder bool) {

	t := initTable(textColour)

	filesList := getFiles(gitFolder)
	totalFiles, extCount := countFiles(filesList)
	total := totalFiles["normal"] + totalFiles["hidden"]
	dirsList := countDirs()
	status := func() map[string]string {
		if gs, ok := getGitStatus(); ok {
			return map[string]string{
				"branch":   fmt.Sprintf("%v,", gs["branch"]),
				"added":    fmt.Sprintf("%v added", gs["added"]),
				"modified": fmt.Sprintf("%v modified", gs["modified"]),
				"deleted":  fmt.Sprintf("%v deleted", gs["deleted"]),
			}
		}
		return map[string]string{
			"branch":   "not a repo",
			"added":    "",
			"modified": "",
			"deleted":  "",
		}
	}

	if printFormat == "long" && len(extCount) > 2 {
		t.AppendRows([]table.Row{
			{fmt.Sprintf("Files%v", separator), fmt.Sprintf("%v regular + %v hidden (%v%% %v, %v%% %v, %v%% %v)",
				totalFiles["normal"], totalFiles["hidden"],
				100*extCount[0].Value/total, extCount[0].Key,
				100*extCount[1].Value/total, extCount[1].Key,
				100*extCount[2].Value/total, extCount[2].Key)},
			{fmt.Sprintf("Dirs%v", separator), fmt.Sprintf("%v   + %v  /  + %v  / /  ",
				dirsList["one"], dirsList["two"], dirsList["three"])},
			{fmt.Sprintf("Git%v", separator), fmt.Sprintf("\uE0A0 %v", fmt.Sprintf("%v %v %v %v",
				status()["branch"], status()["added"], status()["modified"], status()["deleted"]))},
		})
	} else {
		t.AppendRows([]table.Row{
			{fmt.Sprintf("Files%v", separator), fmt.Sprintf("%v + %v", totalFiles["normal"], totalFiles["hidden"])},
			{fmt.Sprintf("Dirs%v", separator), fmt.Sprintf("%v + %v + %v ", dirsList["one"], dirsList["two"], dirsList["three"])},
			{fmt.Sprintf("Git%v", separator), fmt.Sprintf("%v", status()["branch"])},
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
