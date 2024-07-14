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
