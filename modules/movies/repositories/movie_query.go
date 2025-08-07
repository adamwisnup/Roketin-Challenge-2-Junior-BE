package repositories

import (
	"challenge-2/helpers"
	"challenge-2/models"
	"challenge-2/modules/movies"
	"strings"

	"gorm.io/gorm"
)

type MovieQueryRepository struct {
	db *gorm.DB
}

func NewMovieQuery(db *gorm.DB) movies.MovieQueryInterface {
	return &MovieQueryRepository{db: db}
}

func (r *MovieQueryRepository) GetMovies(p helpers.Pagination) ([]models.Movie, int64, error) {
	var movies []models.Movie
	var total int64

	if err := r.db.Model(&models.Movie{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Scopes(helpers.Paginate(p)).Find(&movies).Error; err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}

func (r *MovieQueryRepository) SearchMovies(query string, p helpers.Pagination) ([]models.Movie, int64, error) {
	var movies []models.Movie
	var total int64

	search := "%" + strings.ToLower(query) + "%"
	if err := r.db.Model(&models.Movie{}).
		Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ? OR LOWER(artists) LIKE ? OR LOWER(genres) LIKE ?",
			search, search, search, search).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Scopes(helpers.Paginate(p)).
		Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ? OR LOWER(artists) LIKE ? OR LOWER(genres) LIKE ?",
			search, search, search, search).
		Find(&movies).Error; err != nil {
		return nil, 0, err
	}

	return movies, total, nil
}
