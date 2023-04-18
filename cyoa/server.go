package main

import (
	"fmt"
	"net/http"
	"strings"
)

type StoryStore interface {
	GetChapter(name string) (Chapter, bool)
}

type StoryServer struct {
	StoryStore StoryStore
}

func (s *StoryServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chapterName := strings.TrimPrefix(r.URL.Path, "/")

	chapter, ok := s.StoryStore.GetChapter(chapterName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Chapter not found")
		return
	}

	chapterPage := makeChapterPage(chapter)
	fmt.Fprint(w, chapterPage)
}

func makeChapterPage(chapter Chapter) string {
	var chapterPage strings.Builder

	chapterPage.WriteString(chapter.Title)
	chapterPage.WriteString("\n\n")

	for _, s := range chapter.Paragraphs {
		chapterPage.WriteString(s)
		chapterPage.WriteString("\n\n")
	}

	chapterPage.WriteString("---\n\n")

	for _, o := range chapter.Options {
		chapterPage.WriteString(o.Text)
		fmt.Fprintf(&chapterPage, " <%s>", o.Arc)
		chapterPage.WriteString("\n\n")
	}

	return strings.TrimSuffix(chapterPage.String(), "\n\n")
}
