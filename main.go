package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	_ "go-demo-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
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
// @Produce json
// @Success 200 {array} Movie
// @Router /movies [get]
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")

	// use the line below to serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":9000", r)
}
