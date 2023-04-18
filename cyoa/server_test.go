package main

import (
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

	storyServer := StoryServer{&storyStore}

	t.Run("get intro page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/intro", nil)
		response := httptest.NewRecorder()

		storyServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), `The Little Blue Gopher

Once upon a time, ...

One of his friends once ...

---

That story about ... <new-york>

Gee, those bandits ... <denver>`)
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
