package notes

type SectionsList struct {
	NoteSection []NoteSection `json:"Section"`
}

type NoteSection struct {
	Name  []string            `json:"Name"`
	Notes []map[string]string `json:"Notes"`
}
