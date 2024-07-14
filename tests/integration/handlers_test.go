//go:build integration
// +build integration

package integration_test

import (
	"encoding/json"
	"go-demo-api/internal/api"
	"go-demo-api/internal/db"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetMovies_Integration(t *testing.T) {
	// Initialize the database connection
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Instantiate the repository
	repo := &db.MovieRepository{DB: database}

	// Instantiate the handler struct with the repository
	movieHandler := &api.MovieHandler{Repo: repo}

	// Setup router and server for testing
	r := mux.NewRouter()
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/movies")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	var movies []db.Movie
	err = json.NewDecoder(resp.Body).Decode(&movies)
	if err != nil {
		t.Fatal(err)
	}

	assert.Greater(t, len(movies), 0, "Expected at least one movie")
}

func TestGetMovie_Integration(t *testing.T) {
	// Initialize the database connection
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Instantiate the repository
	repo := &db.MovieRepository{DB: database}

	// Instantiate the handler struct with the repository
	movieHandler := &api.MovieHandler{Repo: repo}

	// Setup router and server for testing
	r := mux.NewRouter()
	r.HandleFunc("/movies/{id}", movieHandler.GetMovie).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/movies/13d7bd2f-732b-465a-bcbe-4b0bc58c3fad")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	var movie db.Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		t.Fatal(err)
	}

	// Log the movie details
	log.Printf("TEst: %+v\n", movie)

	assert.Equal(t, "Inception", movie.Title, "Expected the movie ID to match")
}
