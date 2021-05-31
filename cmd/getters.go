package cmd

import (
	"io/fs"
	"log"
	"os/exec"
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

func countDirs() map[string]int {
	var gr = regexp.MustCompile(`^\.git`)
	dirList := map[string]int{
		"one":   0,
		"two":   0,
		"three": 0,
	}

	err := filepath.WalkDir(".",
		func(path string, d fs.DirEntry, e error) error {
			if e != nil {
				return e
			}
			if d.IsDir() && path != "." && !gr.MatchString(path) {
				depth := strings.Split(path, "/")
				if len(depth) == 1 {
					dirList["one"]++
				} else if len(depth) == 2 {
					dirList["two"]++
				} else {
					dirList["three"]++
				}
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
	return dirList
}

func isGitRepo() bool {
	_, err := exec.Command("bash", "-c", "git rev-parse --is-inside-work-tree 2>/dev/null").Output()
	if err != nil {
		return false
	}
	return true
}

func getGitStatus() map[string]string {
	gitBranch, eb := exec.Command("bash", "-c", "git branch --show-current | tr -d '\n' ").Output()
	if eb != nil {
		log.Fatal(eb)
	}

	modified, em := exec.Command("bash", "-c", "git diff --name-only --diff-filter=M | wc -l | tr -d '\n' | tr -d ' '").Output()
	if em != nil {
		log.Fatal(em)
	}

	added, ea := exec.Command("bash", "-c", "git diff --name-only --diff-filter=A | wc -l | tr -d '\n' | tr -d ' '").Output()
	if ea != nil {
		log.Fatal(ea)
	}

	deleted, ed := exec.Command("bash", "-c", "git diff --name-only --diff-filter=D | wc -l | tr -d '\n' | tr -d ' '").Output()
	if ed != nil {
		log.Fatal(ed)
	}

	return map[string]string{
		"branch":   string(gitBranch),
		"modified": string(modified),
		"added":    string(added),
		"deleted":  string(deleted),
	}
}
