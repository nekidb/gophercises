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

func Render(w io.Writer, chapter Chapter) error {
	templ, err := template.ParseFS(chapterTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "chapter.gohtml", chapter); err != nil {
		return err
	}

	return nil
}
