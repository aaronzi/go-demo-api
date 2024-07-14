package api

import (
	"encoding/json"
	"errors"
	"go-demo-api/internal/db"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie struct to hold movie data
type APIMovie struct {
	db.Movie
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}

type MovieHandler struct {
	Repo db.MovieRepositoryInterface
}

// getMovies godoc
// @Summary Retrieve list of movies
// @Description Get all movies
// @Tags movies
// @Produce json
// @Success 200 {array} APIMovie
// @Router /movies [get]
func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Repo.FindAllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// getMovie godoc
// @Summary Get a movie
// @Description Get a single movie by its ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param   id   path    string     true  "Movie ID"
// @Success 200  {object}  APIMovie
// @Failure 404  {object}  nil  "Movie not found"
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	movie, err := h.Repo.FindMovieByID(id)
	if err != nil {
		// Check if the error is a "not found" error
		if errors.Is(err, db.ErrNotFound) {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if movie == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
