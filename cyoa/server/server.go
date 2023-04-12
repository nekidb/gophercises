package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nekidb/gophercises/cyoa/stories"
)

type StoriesStore interface {
	GetStory(name string) (stories.Story, bool)
}

type StoriesServer struct {
	StoriesStore StoriesStore
}

func (s *StoriesServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storyName := strings.TrimPrefix(r.URL.Path, "/")

	story, ok := s.StoriesStore.GetStory(storyName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Chapter not found")
		return
	}

	storyPage := makeStoryPage(story)
	fmt.Fprint(w, storyPage)
}

func makeStoryPage(story stories.Story) string {
	var storyPage strings.Builder

	storyPage.WriteString(story.Title)
	storyPage.WriteString("\n\n")

	for _, s := range story.Story {
		storyPage.WriteString(s)
		storyPage.WriteString("\n\n")
	}

	storyPage.WriteString("---\n\n")

	for _, o := range story.Options {
		storyPage.WriteString(o.Text)
		fmt.Fprintf(&storyPage, " <%s>", o.Arc)
		storyPage.WriteString("\n\n")
	}

	return strings.TrimSuffix(storyPage.String(), "\n\n")
}
