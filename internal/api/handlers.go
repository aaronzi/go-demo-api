package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie struct to hold movie data
type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     string `json:"year"`
}

// Movies slice to seed movie data
var movies = []Movie{
	{ID: "1", Title: "Inception", Director: "Christopher Nolan", Year: "2010"},
	{ID: "2", Title: "The Matrix", Director: "Lana Wachowski, Lilly Wachowski", Year: "1999"},
	{ID: "3", Title: "Interstellar", Director: "Christopher Nolan", Year: "2014"},
}

// getMovies godoc
// @Summary Retrieve list of movies
// @Description Get all movies
// @Tags movies
// @Produce json
// @Success 200 {array} Movie
// @Router /movies [get]
func GetMovies(w http.ResponseWriter, r *http.Request) {
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
// @Success 200  {object}  Movie
// @Failure 404  {object}  nil  "Movie not found"
// @Router /movies/{id} [get]
func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, movie := range movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.NotFound(w, r)
}
