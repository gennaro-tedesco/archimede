package cmd

import (
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func getFiles(gitFolder bool) []string {
	var filesList []string
	var gr = regexp.MustCompile(`^\.git/`)

	err := filepath.WalkDir(".",
		func(path string, d fs.DirEntry, e error) error {
			if e != nil {
				return e
			}
			if !d.IsDir() {
				if gitFolder {
					filesList = append(filesList, path)
				} else if !gr.MatchString(path) {
					filesList = append(filesList, path)
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return filesList
}

func countFiles(filesList []string) (map[string]int, []kv) {
	var normalFiles []string
	var hiddenFiles []string
	var hr = regexp.MustCompile(`^\.`)
	var extensions []string

	for _, f := range filesList {
		if hr.MatchString(f) {
			hiddenFiles = append(hiddenFiles, f)
		} else {
			normalFiles = append(normalFiles, f)
		}
		ext := strings.Split(f, ".")
		extensions = append(extensions, ext[len(ext)-1])
	}

	totalFiles := map[string]int{
		"normal": len(normalFiles),
		"hidden": len(hiddenFiles),
	}

	extCount := make(map[string]int)
	for _, ex := range extensions {
		extCount[ex]++
	}

	return totalFiles, sortExtensions(extCount)
}

func sortExtensions(extCount map[string]int) []kv {

	var ss []kv
	for k, v := range extCount {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	return ss
}
