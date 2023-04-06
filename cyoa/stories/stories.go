package stories

import (
	"encoding/json"
	"io/fs"
)

type Story struct {
}

type Stories map[string]Story

func GetStoriesFromFile(fileSystem fs.FS, fileName string) (Stories, error) {
	fileData, err := fs.ReadFile(fileSystem, fileName)
	if err != nil {
		return nil, err
	}

	stories := make(map[string]Story)
	if err := json.Unmarshal(fileData, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}
