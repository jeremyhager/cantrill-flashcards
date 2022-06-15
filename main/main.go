package main

type SectionsList struct {
	NoteSection []NoteSection
}

type NoteSection struct {
	Name  []string            `json:"Name"`
	Notes []map[string]string `json:"Notes"`
}

func main() {

}
