package main

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		chapter = Chapter{
			Title: "The Little Blue Gopher",
			Paragraphs: []string{
				"Once upon a time ...",
				"One of his friends ...",
			},
			Options: []Option{
				{Text: "That story about ...", Arc: "new-york"},
				{Text: "Gee, those bandits ...", Arc: "denver"},
			},
		}
	)

	t.Run("it renders chapter to html page", func(t *testing.T) {
		buf := bytes.Buffer{}
		chapterRenderer, err := NewChapterRenderer()
		if err != nil {
			t.Fatal(err)
		}

		if err := chapterRenderer.Render(&buf, chapter); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		chapter = Chapter{
			Title: "The Little Blue Gopher",
			Paragraphs: []string{
				"Once upon a time ...",
				"One of his friends ...",
			},
			Options: []Option{
				{Text: "That story about ...", Arc: "new-york"},
				{Text: "Gee, those bandits ...", Arc: "denver"},
			},
		}
	)

	chapterRenderer, err := NewChapterRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chapterRenderer.Render(io.Discard, chapter)
	}
}
