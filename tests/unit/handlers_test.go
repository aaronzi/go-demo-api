//go:build unit
// +build unit

package unit_test

import (
	"encoding/json"
	"go-demo-api/internal/api"
	"go-demo-api/internal/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockRepository struct{}

func (m *MockRepository) FindAllMovies() ([]db.Movie, error) {
	return []db.Movie{{ID: "1", Title: "Mock Movie", Director: "Mock Director", Year: 2010}}, nil
}

func (m *MockRepository) FindMovieByID(id string) (*db.Movie, error) {
	// Return a movie or an error based on your test scenario
	if id == "1" {
		return &db.Movie{ID: "1", Title: "Mock Movie", Director: "Mock Director", Year: 2010}, nil
	}
	return nil, db.ErrNotFound
}

// Tests the GetMovies handler
func TestGetMovies_Unit(t *testing.T) {
	mockRepo := &MockRepository{}
	handler := api.MovieHandler{Repo: mockRepo}

	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies", handler.GetMovies)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")

	var movies []api.APIMovie
	err = json.Unmarshal(rr.Body.Bytes(), &movies)
	if err != nil {
		t.Fatal(err)
	}

	assert.Greater(t, len(movies), 0, "Expected at least one movie")
}

// Tests the GetMovie handler with a valid movie ID.
func TestGetMovie_ValidID(t *testing.T) {
	mockRepo := &MockRepository{}
	handler := api.MovieHandler{Repo: mockRepo}

	req, err := http.NewRequest("GET", "/movies/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", handler.GetMovie)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")

	var movie api.APIMovie
	err = json.Unmarshal(rr.Body.Bytes(), &movie)
	if err != nil {
		t.Fatal(err)
	}

	expected := api.APIMovie{
		Movie: db.Movie{
			ID:       "1",
			Title:    "Mock Movie",
			Director: "Mock Director",
			Year:     2010,
		},
	}

	assert.Equal(t, expected, movie, "Expected the movie ID to match")
}

// Tests the GetMovie handler with an invalid movie ID.
func TestGetMovie_InvalidID(t *testing.T) {
	mockRepo := &MockRepository{}
	handler := api.MovieHandler{Repo: mockRepo}

	req, err := http.NewRequest("GET", "/movies/nonExistentID", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", handler.GetMovie)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Expected status code 404")
}
