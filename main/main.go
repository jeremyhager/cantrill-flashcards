package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type SectionsList struct {
	NoteSection []NoteSection `json:"Section"`
}

type NoteSection struct {
	Name  []string            `json:"Name"`
	Notes []map[string]string `json:"Notes"`
}

func main() {

	sections := &SectionsList{}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == "main.go" {
			continue
		} else {
			currentFile, err := ioutil.ReadFile(file.Name())
			if err != nil {
				log.Fatal(err)
			}

			_ = json.Unmarshal(currentFile, &sections)

			for k := range sections.NoteSection[0].Notes {

			}

			fmt.Printf("%+v", sections.NoteSection[0].Notes)
		}
	}
}

func keysList(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0

	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
