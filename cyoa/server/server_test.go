package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nekidb/gophercises/cyoa/server"
	"github.com/nekidb/gophercises/cyoa/stories"
)

func TestStoriesServer(t *testing.T) {
	storiesStore := StubStoriesStore{
		stories.Stories{
			"intro": stories.Story{
				Title: "The Little Blue Gopher",
				Story: []string{
					"Once upon a time, ...",
					"One of his friends once ...",
				},
				Options: []stories.Option{
					{Text: "That story about ...", Arc: "new-york"},
					{Text: "Gee, those bandits ...", Arc: "denver"},
				},
			},
		},
	}

	storiesServer := server.StoriesServer{&storiesStore}

	t.Run("get intro page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/intro", nil)
		response := httptest.NewRecorder()

		storiesServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), `The Little Blue Gopher

Once upon a time, ...

One of his friends once ...

---

That story about ... <new-york>

Gee, those bandits ... <denver>`)
	})
	t.Run("get not existing story", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/badstory", nil)
		response := httptest.NewRecorder()

		storiesServer.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseBody(t, response.Body.String(), `Chapter not found`)
	})
}

type StubStoriesStore struct {
	stories stories.Stories
}

func (s *StubStoriesStore) GetStory(name string) (stories.Story, bool) {
	story, ok := s.stories[name]
	if !ok {
		return stories.Story{}, false
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
