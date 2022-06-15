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
			for q, a := range sections.NoteSection[0].Notes {
				fmt.Printf("question: %v\n", q)
				fmt.Printf("answer: %v\n", a)
			}
		}
	}
}
