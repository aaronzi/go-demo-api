package api

import (
	"encoding/json"
	"net/http"
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

// GetMovies retrieves list of movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
