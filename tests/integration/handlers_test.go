//go:build integration
// +build integration

package integration_test

import (
	"encoding/json"
	"go-demo-api/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetMovies_Integration(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/movies", api.GetMovies).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/movies")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	var movies []api.Movie
	err = json.NewDecoder(resp.Body).Decode(&movies)
	if err != nil {
		t.Fatal(err)
	}

	assert.Greater(t, len(movies), 0, "Expected at least one movie")
}

func TestGetMovie_Integration(t *testing.T) {
	r := mux.NewRouter()
	// Assuming api.GetMovie is the handler for getting a single movie by ID
	r.HandleFunc("/movies/{id}", api.GetMovie).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	// Replace "someValidMovieID" with a valid movie ID that exists in your data source
	resp, err := http.Get(ts.URL + "/movies/1")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	var movie api.Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		t.Fatal(err)
	}

	// Replace "someValidMovieID" with the same valid movie ID used in the request
	assert.Equal(t, "1", movie.ID, "Expected the movie ID to match")
}
