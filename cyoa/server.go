package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type StoryStore interface {
	GetChapter(name string) (Chapter, bool)
}

type Renderer interface {
	Render(io.Writer, Chapter) error
}

type StoryServer struct {
	StoryStore StoryStore
	Renderer   Renderer
}

func (s *StoryServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chapterName := strings.TrimPrefix(r.URL.Path, "/")

	chapter, ok := s.StoryStore.GetChapter(chapterName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Chapter not found")
		return
	}

	if err := s.Renderer.Render(w, chapter); err != nil {
		panic(err)
	}
}
