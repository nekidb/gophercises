package stories_test

import (
	"testing"
	"testing/fstest"

	"github.com/nekidb/gophercises/cyoa/stories"
)

func TestGetStoriesFromFile(t *testing.T) {
	json := `{"story1":{},"story2":{}}`

	fileName := "stories.json"
	fs := fstest.MapFS{
		fileName: {Data: []byte(json)},
	}

	stories, err := stories.GetStoriesFromFile(fs, fileName)
	if err != nil {
		t.Fatal(err)
	}

	got := len(stories)
	want := 2

	if got != want {
		t.Errorf("wanted %d stories, but got %d", want, got)
	}
}
