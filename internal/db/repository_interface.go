package db

// MovieRepositoryInterface defines the methods that a movie repository must have.
type MovieRepositoryInterface interface {
	FindAllMovies() ([]Movie, error)
	FindMovieByID(id string) (*Movie, error)
}
