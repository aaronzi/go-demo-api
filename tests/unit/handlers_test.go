//go:build unit
// +build unit

package unit_test

import (
	"encoding/json"
	"go-demo-api/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetMovies_Unit(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetMovies)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")

	var movies []api.Movie
	err = json.Unmarshal(rr.Body.Bytes(), &movies)
	if err != nil {
		t.Fatal(err)
	}

	assert.Greater(t, len(movies), 0, "Expected at least one movie")
}

func TestGetMovie_ValidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": "1", // Replace with an actual valid ID
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetMovie)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")

	var movie api.Movie
	err = json.Unmarshal(rr.Body.Bytes(), &movie)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "1", movie.ID, "Expected the movie ID to match")
}

func TestGetMovie_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/movies/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": "nonExistentID",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetMovie)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Expected status code 404")
}
