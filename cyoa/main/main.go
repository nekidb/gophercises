package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nekidb/gophercises/cyoa/server"
	"github.com/nekidb/gophercises/cyoa/stories"
)

type InMemoryStoriesStore struct {
	stories stories.Stories
}

func (i *InMemoryStoriesStore) GetStory(name string) (stories.Story, bool) {
	story, ok := i.stories[name]
	if !ok {
		return stories.Story{}, false
	}
	return story, true
}

func main() {
	storiesFromFile, err := stories.GetStoriesFromFile(os.DirFS("."), "stories.json")
	if err != nil {
		panic(err)
	}
	storiesStore := InMemoryStoriesStore{storiesFromFile}
	storiesServer := server.StoriesServer{&storiesStore}
	log.Fatal(http.ListenAndServe(":8080", &storiesServer))
}
