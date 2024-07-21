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

// corsMiddleware sets up the CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                                                   // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")                                             // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization") // Allowed headers

		// Check if the request is for the OPTIONS method (pre-flight request)
		if r.Method == "OPTIONS" {
			// Respond with OK status for pre-flight requests
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		next.ServeHTTP(w, r)
	})
}

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

	r.Use(corsMiddleware) // Use the CORS middleware

	// Use the methods of movieHandler as HTTP handlers
	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", movieHandler.GetMovie).Methods("GET")
	r.HandleFunc("/health", movieHandler.HealthCheckHandler).Methods("GET")

	// Use the methods of userHandler as HTTP handlers
	r.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/verify", verifyHandler.VerifyUser).Methods("GET")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	// Serve Swagger UI
	r.PathPrefix("/swagger-ui/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":9000", r)
}
