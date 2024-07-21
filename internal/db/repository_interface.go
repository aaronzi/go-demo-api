package db

// MovieRepositoryInterface defines the methods that a movie repository must have.
type MovieRepositoryInterface interface {
	FindAllMovies() ([]Movie, error)
	FindMovieByID(id string) (*Movie, error)
}

type UserRepositoryInterface interface {
	RegisterUser(username string, email string, password string) error
	LoginUser(username string, password string) (string, error)
}

type VerificationRepositoryInterface interface {
	CreateVerification(email string) (string, error)
	Verify(email string, verificationCode string) (string, error)
}
