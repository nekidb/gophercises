package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStoriesServer(t *testing.T) {
	storyStore := StubStoryStore{
		Story{
			"intro": Chapter{
				Title: "The Little Blue Gopher",
				Paragraphs: []string{
					"Once upon a time, ...",
					"One of his friends once ...",
				},
				Options: []Option{
					{Text: "That story about ...", Arc: "new-york"},
					{Text: "Gee, those bandits ...", Arc: "denver"},
				},
			},
		},
	}

	renderer := StubRenderer{}

	storyServer := StoryServer{StoryStore: &storyStore, Renderer: &renderer}

	t.Run("it call Render when chapter exists", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/intro", nil)
		response := httptest.NewRecorder()

		storyServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		if renderer.Calls != 1 {
			t.Errorf("did not get right render calls, got %d, want %d", renderer.Calls, 1)
		}
	})
	t.Run("get not existing chapter", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/badchapter", nil)
		response := httptest.NewRecorder()

		storyServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseBody(t, response.Body.String(), `Chapter not found`)
	})
}

type StubStoryStore struct {
	story Story
}

func (s *StubStoryStore) GetChapter(name string) (Chapter, bool) {
	story, ok := s.story[name]
	if !ok {
		return Chapter{}, false
	}
	return story, true
}

type StubRenderer struct {
	Calls int
}

func (s *StubRenderer) Render(w io.Writer, chapter Chapter) error {
	s.Calls++

	return nil
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("gid not get correct status, got %d, want %d", got, want)
	}
}
