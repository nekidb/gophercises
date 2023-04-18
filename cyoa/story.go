package main

import (
	"encoding/json"
	"io/fs"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func GetStoryFromFile(fileSystem fs.FS, fileName string) (Story, error) {
	fileData, err := getDataFromFile(fileSystem, fileName)
	if err != nil {
		return nil, err
	}

	story, err := parseJSON(fileData)
	if err != nil {
		return nil, err
	}

	return story, nil
}

func getDataFromFile(fileSystem fs.FS, fileName string) ([]byte, error) {
	return fs.ReadFile(fileSystem, fileName)
}

func parseJSON(jsonData []byte) (Story, error) {
	story := make(map[string]Chapter)
	if err := json.Unmarshal(jsonData, &story); err != nil {
		return nil, err
	}
	return story, nil
}
