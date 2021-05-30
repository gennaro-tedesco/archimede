package cmd

import (
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
)

func getFiles() []string {
	var filesList []string
	var gr = regexp.MustCompile(`^\.git/`)

	err := filepath.WalkDir(".",
		func(path string, d fs.DirEntry, e error) error {
			if e != nil {
				return e
			}
			if !d.IsDir() && !gr.MatchString(path) {
				filesList = append(filesList, path)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return filesList
}

func parseFiles(filesList []string) map[string]int {
	var normalFiles []string
	var hiddenFiles []string
	var hr = regexp.MustCompile(`^\.`)

	for _, f := range filesList {
		if hr.MatchString(f) {
			hiddenFiles = append(hiddenFiles, f)
		} else {
			normalFiles = append(normalFiles, f)
		}
	}

	counts := map[string]int{
		"normal": len(normalFiles),
		"hidden": len(hiddenFiles),
	}

	return counts
}
