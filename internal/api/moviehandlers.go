package api

import (
	"encoding/json"
	"errors"
	"go-demo-api/internal/db"

	auth "go-demo-api/internal/auth"
	utils "go-demo-api/internal/util"
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
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /movies [get]
func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	// Check for auth token in the Authorization header
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return
	}
	secret, fileReadError := utils.ReadFile("/workspace/privatekey.txt")

	if fileReadError != nil {
		http.Error(w, "Error reading token", http.StatusInternalServerError)
	}
	// Token validation
	isValid, err := auth.IsTokenValid(authToken, secret)
	if err != nil || !isValid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

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
// @Failure 401  {object}  nil  "Unauthorized"
// @Failure 404  {object}  nil  "Movie not found"
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	// Check for auth token in the Authorization header
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return
	}
	secret, fileReadError := utils.ReadFile("/workspace/privatekey.txt")

	if fileReadError != nil {
		http.Error(w, "Error reading token", http.StatusInternalServerError)
	}
	// Token validation
	isValid, err := auth.IsTokenValid(authToken, secret)
	if err != nil || !isValid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

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

// HealthCheckHandler responds to health check requests
// @Summary Health Check
// @Description Responds with OK if the service is up and running
// @Tags health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /health [get]
func (h *MovieHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
