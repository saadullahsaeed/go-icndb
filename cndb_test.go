package chucknorris

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRandomJoke(t *testing.T) {
	tests := []struct {
		req      *CNDBRequest
		jokeID   int
		jokeText string
	}{
		{nil, 1, "Joke"},
	}

	for i, tt := range tests {
		mockResp := fmt.Sprintf(
			`{ "type": "success", "value": { "id": %d, "joke": "%s", "categories": [] } }`,
			tt.jokeID,
			tt.jokeText,
		)
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, mockResp)
		}))
		baseURL = ts.URL

		joke, err := GetRandomJoke(nil)
		if err != nil {
			t.Fatal(err)
		}

		if joke.ID != tt.jokeID {
			t.Errorf("Test - %d, Expected - %s , Got - %s", i, tt.jokeID, joke.ID)
			continue
		}

		if joke.ID != tt.jokeID || joke.String() != tt.jokeText {
			t.Errorf("Test - %d, Expected - %s , Got - %s", i, tt.jokeText, joke)
			continue
		}
	}
}

func TestGetJokeByID(t *testing.T) {
	mockResp := `{ "type": "success", "value": { "id": 1, "joke": "test", "categories": [] } }`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, mockResp)
	}))
	baseURL = ts.URL

	joke, err := GetJokeByID(1, nil)
	if err != nil {
		t.Fatal(err)
	}

	if joke.String() != "test" {
		t.Errorf("Expected - test , Got - %s", joke)
	}
}
