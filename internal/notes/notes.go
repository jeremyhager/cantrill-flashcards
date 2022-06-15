package notes

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
)

type Inotes interface {
}

type Client struct {
}

func (c *Client) OpenFile(fileName string) ([]byte, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil

}

func (c *Client) parseNotes(data []byte) *SectionsList {
	openMap := &SectionsList{}
	fmt.Println(openMap)
	_ = json.Unmarshal(data, &openMap)

	return openMap
}

func (c *Client) ParseSection(fileName string) (*SectionsList, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	section := c.parseNotes(file)

	return section, nil
}

func (c *Client) OpenFiles(directory string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (c *Client) CreateQuestion(files []fs.FileInfo) (string, error) {
	for _, file := range files {
		sections, err := c.ParseSection(file.Name())
		if err != nil {
			return "", err
		}

	}
}

func (c *Client) randomize() (string, error) {

}
