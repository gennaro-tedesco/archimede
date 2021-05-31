package cmd

import (
	"io/fs"
	"log"
	"os"
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

func getCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
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
	_, err := exec.Command("sh", "-c", "git rev-parse --is-inside-work-tree 2>/dev/null").Output()
	if err != nil {
		return false
	}
	return true
}

func getGitStatus() map[string]string {
	gitBranch, eb := exec.Command("sh", "-c", "git branch --show-current | tr -d '\n' ").Output()
	if eb != nil {
		log.Fatal(eb)
	}

	modified, em := exec.Command("sh", "-c", "git diff --name-only --diff-filter=M | wc -l | tr -d '\n' | tr -d ' '").Output()
	if em != nil {
		log.Fatal(em)
	}

	staged, ed := exec.Command("sh", "-c", "git diff --name-only --staged | wc -l | tr -d '\n' | tr -d ' '").Output()
	if ed != nil {
		log.Fatal(ed)
	}

	return map[string]string{
		"branch":   string(gitBranch),
		"modified": string(modified),
		"staged":   string(staged),
	}
}

func getDiskUsage() string {
	du, err := exec.Command("sh", "-c", "du -sh . | cut -f1 | tr -d '\n' | tr -d ' '").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(du)
}
