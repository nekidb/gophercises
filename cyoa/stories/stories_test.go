package stories_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/nekidb/gophercises/cyoa/stories"
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

	fileName := "stories.json"
	fs := fstest.MapFS{
		fileName: {Data: []byte(jsonData)},
	}

	storiesList, err := stories.GetStoriesFromFile(fs, fileName)
	if err != nil {
		t.Fatal(err)
	}

	assertStory(t, storiesList["intro"], stories.Story{
		Title: "The Little Blue Gopher",
		Story: []string{"This is intro", "This is end of intro"},
		Options: []stories.Option{
			stories.Option{"This is option's text", "new-york"},
		},
	})
}

func assertStory(t *testing.T, got, want stories.Story) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
