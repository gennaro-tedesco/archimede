package cmd

import (
	"io/ioutil"
	"log"
)

func getFiles() []string {
	var filesList []string

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			filesList = append(filesList, f.Name())
		}
	}
	return filesList
}
