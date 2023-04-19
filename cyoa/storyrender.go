package main

import (
	"html/template"
	"io"
)

func Render(w io.Writer, chapter Chapter) error {

	chapterTemplate := `<h1>{{ .Title }}</h1>

{{ range .Paragraphs }}<p>{{ . }}</p>
{{end}}
<ul>{{ range .Options }}
<li><a href="/{{ .Arc }}">{{ .Text }}</a></li>{{end}}
</ul>
`

	templ, err := template.New("chapter").Parse(chapterTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, chapter); err != nil {
		return err
	}

	return nil
}
