package services

import (
	"challenge-2/helpers"
	"challenge-2/models"
	"challenge-2/modules/movies"
	"errors"
)

type MovieService struct {
	rq movies.MovieQueryInterface
	rc movies.MovieCommandInterface
}

func NewMovieService(rq movies.MovieQueryInterface, rc movies.MovieCommandInterface) *MovieService {
	return &MovieService{
		rq: rq,
		rc: rc,
	}
}

func (s *MovieService) GetMovies(p helpers.Pagination) ([]models.Movie, int64, error) {
	return s.rq.GetMovies(p)
}

func (s *MovieService) CreateMovie(movie models.Movie) (models.Movie, error) {
	if movie.Title == "" || movie.Description == "" || movie.Duration <= 0 {
		return models.Movie{}, errors.New("invalid movie data")
	}
	return s.rc.CreateMovie(movie)
}

func (s *MovieService) UpdateMovie(id int, data map[string]interface{}) (models.Movie, error) {
	if len(data) == 0 {
		return models.Movie{}, errors.New("no data to update")
	}

	updatedMovie, err := s.rc.UpdateMovie(id, data)
	if err != nil {
		return models.Movie{}, err
	}

	return updatedMovie, nil
}

func (s *MovieService) SearchMovies(query string, p helpers.Pagination) ([]models.Movie, int64, error) {
	if query == "" {
		return nil, 0, errors.New("search query cannot be empty")
	}

	return s.rq.SearchMovies(query, p)
}

func (s *MovieService) UploadMovies(movies []models.Movie) error {
	if len(movies) == 0 {
		return errors.New("no movies to upload")
	}

	return s.rc.BulkInsertMovie(movies)
}