package stories

import (
	"encoding/json"
	"io/fs"
)

type Story struct {
	Title string
}

type Stories map[string]Story

func GetStoriesFromFile(fileSystem fs.FS, fileName string) (Stories, error) {
	fileData, err := getDataFromFile(fileSystem, fileName)
	if err != nil {
		return nil, err
	}

	stories, err := parseJSON(fileData)
	if err != nil {
		return nil, err
	}

	return stories, nil
}

func getDataFromFile(fileSystem fs.FS, fileName string) ([]byte, error) {
	return fs.ReadFile(fileSystem, fileName)
}

func parseJSON(jsonData []byte) (Stories, error) {
	stories := make(map[string]Story)
	if err := json.Unmarshal(jsonData, &stories); err != nil {
		return nil, err
	}
	return stories, nil
}
