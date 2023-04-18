package main

import (
	"log"
	"net/http"
	"os"
)

type InMemoryStoriesStore struct {
	stories Story
}

func (i *InMemoryStoriesStore) GetChapter(name string) (Chapter, bool) {
	story, ok := i.stories[name]
	if !ok {
		return Chapter{}, false
	}
	return story, true
}

func main() {
	storiesFromFile, err := GetStoryFromFile(os.DirFS("."), "stories.json")
	if err != nil {
		panic(err)
	}
	storiesStore := InMemoryStoriesStore{storiesFromFile}
	storiesServer := StoryServer{&storiesStore}
	log.Fatal(http.ListenAndServe(":8080", &storiesServer))
}
