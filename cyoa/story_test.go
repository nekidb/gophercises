package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestGetStoriesFromFile(t *testing.T) {
	const jsonData = `
{
	"intro": {
		"title": "The Little Blue Gopher",
		"story": [
			"This is intro",
			"This is end of intro"
		],
		"options": [
			{
				"text": "This is option's text",
				"arc": "new-york"
			}
		]
	},
	"new-york": {
		"title": "Visiting New York"
	}
}`

	fileName := "story.json"
	fs := fstest.MapFS{
		fileName: {Data: []byte(jsonData)},
	}

	story, err := GetStoryFromFile(fs, fileName)
	if err != nil {
		t.Fatal(err)
	}

	assertChapter(t, story["intro"], Chapter{
		Title:      "The Little Blue Gopher",
		Paragraphs: []string{"This is intro", "This is end of intro"},
		Options: []Option{
			Option{"This is option's text", "new-york"},
		},
	})
}

func assertChapter(t *testing.T, got, want Chapter) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
