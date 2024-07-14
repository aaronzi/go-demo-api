package db

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("movie not found")

type MovieRepository struct {
	DB *sql.DB
}

type Movie struct {
	ID       string
	Title    string
	Director string
	Year     int
}

func (repo *MovieRepository) FindAllMovies() ([]Movie, error) {
	rows, err := repo.DB.Query("SELECT id, title, director, year FROM Movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var m Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.Director, &m.Year); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func (repo *MovieRepository) FindMovieByID(id string) (*Movie, error) {
	var m Movie
	err := repo.DB.QueryRow("SELECT id, title, director, year FROM Movies WHERE id = ?", id).Scan(&m.ID, &m.Title, &m.Director, &m.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &m, nil
}
