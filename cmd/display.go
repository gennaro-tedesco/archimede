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
	short bool,
	git bool,
	excludeDir string,
	excludeFile string,
	textColour string,
	delimiter string) {

	t := createTable(textColour)
	displayCwd(t, delimiter)
	displayFiles(t, short, git, excludeDir, excludeFile, delimiter)
	displayDirs(t, short, excludeDir, delimiter)
	displayGit(t, short, delimiter)
	displayDiskUsage(t, delimiter)
	t.AppendSeparator()
	t.Render()
}

func createTable(textColour string) table.Writer {
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
	t.Style().Box.PaddingLeft = "  "
	t.Style().Box.PaddingRight = "  "
	t.Style().Box.BottomLeft = "╰"
	t.Style().Box.TopLeft = "╭"
	t.Style().Box.TopRight = "╮"
	t.Style().Box.BottomRight = "╯"
	return t
}

func displayCwd(t table.Writer, delimiter string) {
	t.AppendRow(table.Row{fmt.Sprintf("Path%v", delimiter), fmt.Sprint(getCwd())})
}

func displayFiles(t table.Writer, short bool, git bool, excludeDir string, excludeFile string, delimiter string) {
	filesList := getFiles(git, excludeDir, excludeFile)
	totalFiles, extCount := countFiles(filesList)
	total := totalFiles["normal"] + totalFiles["hidden"]

	if !short && len(extCount) > 2 {
		t.AppendRow(table.Row{fmt.Sprintf("Files%v", delimiter),
			fmt.Sprintf("%v regular + %v hidden (%v%% %v, %v%% %v, %v%% %v)",
				totalFiles["normal"], totalFiles["hidden"],
				100*extCount[0].Value/total, extCount[0].Key,
				100*extCount[1].Value/total, extCount[1].Key,
				100*extCount[2].Value/total, extCount[2].Key)},
		)
	} else {
		t.AppendRow(table.Row{fmt.Sprintf("Files%v", delimiter),
			fmt.Sprintf("%v + %v", totalFiles["normal"], totalFiles["hidden"])},
		)
	}
}

func displayDirs(t table.Writer, short bool, excludeDir string, delimiter string) {
	dirsList := countDirs(excludeDir)
	if !short {
		t.AppendRow(table.Row{fmt.Sprintf("Dirs%v", delimiter),
			fmt.Sprintf("%v   + %v  /  + %v  / /  ",
				dirsList["one"], dirsList["two"], dirsList["three"])},
		)
	} else {
		t.AppendRow(table.Row{fmt.Sprintf("Dirs%v", delimiter),
			fmt.Sprintf("%v + %v + %v ", dirsList["one"], dirsList["two"], dirsList["three"])},
		)
	}
}

func displayGit(t table.Writer, short bool, delimiter string) {
	if !short {
		if isGitRepo() {
			t.AppendRow(table.Row{fmt.Sprintf("Git%v", delimiter),
				fmt.Sprintf("\uE0A0 %v (%v staged %v modified)",
					getGitStatus()["branch"], getGitStatus()["staged"], getGitStatus()["modified"])},
			)
		} else {
			t.AppendRow(table.Row{fmt.Sprintf("Git%v", delimiter),
				fmt.Sprint("not a git repo")},
			)
		}
	} else {
		if isGitRepo() {
			t.AppendRow(table.Row{fmt.Sprintf("Git%v", delimiter),
				fmt.Sprintf("%v (+%v !%v)",
					getGitStatus()["branch"], getGitStatus()["staged"], getGitStatus()["modified"])},
			)
		} else {
			t.AppendRow(table.Row{fmt.Sprintf("Git%v", delimiter),
				fmt.Sprint("no")},
			)
		}
	}
}

func displayDiskUsage(t table.Writer, delimiter string) {
	t.AppendRow(table.Row{fmt.Sprintf("Space%v", delimiter),
		fmt.Sprintf("%v", getDiskUsage())},
	)
}
