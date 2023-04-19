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
	templ, err := template.ParseFS(chapterTemplates, "templates/chapter.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, chapter); err != nil {
		return err
	}

	return nil
}
