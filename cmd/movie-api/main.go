package main

import (
	_ "go-demo-api/docs"
	"go-demo-api/internal/api"
	"go-demo-api/internal/db"
	user "go-demo-api/internal/db/user"
	verification "go-demo-api/internal/db/verification"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Initialize the database connection
	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Instantiate the repository
	movieRepo := &db.MovieRepository{DB: database}
	verifyRepo := &verification.EmailVerificationRepository{VerificationRepository: &verification.VerificationRepository{DB: database}}
	userRepo := &user.UserRepository{DB: database, VerificationRepository: verifyRepo}

	// Instantiate the handler struct with the repository
	movieHandler := &api.MovieHandler{Repo: movieRepo}
	userHandler := &api.UserHandler{Repo: userRepo}
	verifyHandler := &api.VerificationHandler{VerificationRepository: verifyRepo}

	r := mux.NewRouter()

	// Use the methods of movieHandler as HTTP handlers
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", movieHandler.GetMovie).Methods("GET")
	r.HandleFunc("/health", movieHandler.HealthCheckHandler).Methods("GET")

	// Use the methods of userHandler as HTTP handlers
	r.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/verify", verifyHandler.VerifyUser).Methods("GET")

	// Serve Swagger UI
	r.PathPrefix("/swagger-ui/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":9000", r)
}
