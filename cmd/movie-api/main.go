package main

import (
	_ "go-demo-api/docs"
	"go-demo-api/internal/api"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", api.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", api.GetMovie).Methods("GET")

	// Serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":9000", r)
}
