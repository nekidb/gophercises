package main

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	t.Run("it renders chapter to html page", func(t *testing.T) {
		buf := bytes.Buffer{}
		chapter := Chapter{
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

		if err := Render(&buf, chapter); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
		//
		// 		got := buf.String()
		// 		want := `<h1>The Little Blue Gopher</h1>
		// <p>Once upon a time ...</p>
		// <p>One of his friends ...</p>
		// <ul>
		// <li><a href="/new-york">That story about ...</a></li>
		// <li><a href="/denver">Gee, those bandits ...</a></li>
		// </ul>
		// `
		//
		// 		if got != want {
		// 			t.Errorf("got %q, want %q", got, want)
		// 		}
	})
}
