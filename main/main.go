package main

import "io/ioutil"

type SectionsList struct {
	NoteSection []NoteSection
}

type NoteSection struct {
	Name  []string            `json:"Name"`
	Notes []map[string]string `json:"Notes"`
}

func main() {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {

	}
}
