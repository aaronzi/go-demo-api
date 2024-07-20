package db

// MovieRepositoryInterface defines the methods that a movie repository must have.
type MovieRepositoryInterface interface {
	FindAllMovies() ([]Movie, error)
	FindMovieByID(id string) (*Movie, error)
}

type UserRepositoryInterface interface {
	RegisterUser(username string, email string, password string) error
	CheckUser(username string, password string) error
}

type VerificationRepositoryInterface interface {
	CreateVerification(email string) error
	Verify(email string, verificationCode string) (string, error)
}
