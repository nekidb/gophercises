package main

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	chapterTemplates embed.FS
)

type ChapterRenderer struct {
	templ *template.Template
}

func NewChapterRenderer() (*ChapterRenderer, error) {
	templ, err := template.ParseFS(chapterTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &ChapterRenderer{templ: templ}, nil
}

func (r *ChapterRenderer) Render(w io.Writer, chapter Chapter) error {
	if err := r.templ.ExecuteTemplate(w, "chapter.gohtml", chapter); err != nil {
		return err
	}

	return nil
}
