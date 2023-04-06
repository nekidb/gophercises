package stories_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/nekidb/gophercises/cyoa/stories"
)

func TestGetStoriesFromFile(t *testing.T) {
	jsonData := `
{
	"intro":{"title": "The Little Blue Gopher"},
	"new-york":{"title": "Visiting New York"}
}`

	fileName := "stories.json"
	fs := fstest.MapFS{
		fileName: {Data: []byte(jsonData)},
	}

	storiesList, err := stories.GetStoriesFromFile(fs, fileName)
	if err != nil {
		t.Fatal(err)
	}

	got := storiesList["intro"]
	want := stories.Story{Title: "The Little Blue Gopher"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
