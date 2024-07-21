//go:build integration
// +build integration

package integration_test

import (
	"encoding/json"
	"go-demo-api/internal/api"
	"go-demo-api/internal/db"
	testUtils "go-demo-api/tests"
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
	movieHandler := &api.MovieHandler{Repo: repo, IsTest: false}

	// Setup router and server for testing
	r := mux.NewRouter()
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", ts.URL+"/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the Authorization header
	token, tokenError := testUtils.GenerateToken()
	if tokenError != nil {
		t.Fatal(tokenError)
	}
	req.Header.Add("Authorization", token)

	log.Printf("Request: %+v\n", req)

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Response: %+v\n", resp)

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
	movieHandler := &api.MovieHandler{Repo: repo, IsTest: false}

	// Setup router and server for testing
	r := mux.NewRouter()
	r.HandleFunc("/movies/{id}", movieHandler.GetMovie).Methods("GET")

	ts := httptest.NewServer(r)
	defer ts.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", ts.URL+"/movies/13d7bd2f-732b-465a-bcbe-4b0bc58c3fad", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the Authorization header
	token, tokenError := testUtils.GenerateToken()
	if tokenError != nil {
		t.Fatal(tokenError)
	}
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
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
