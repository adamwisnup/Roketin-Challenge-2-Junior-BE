package repositories

import (
	"challenge-2/models"
	"challenge-2/modules/movies"

	"gorm.io/gorm"
)

type MovieCommandRepository struct {
	db *gorm.DB
}

func NewMovieCommand(db *gorm.DB) movies.MovieCommandInterface {
	return &MovieCommandRepository{db: db}
}

func (r *MovieCommandRepository) CreateMovie(movie models.Movie) (models.Movie, error){
	if err := r.db.Create(&movie).Error; err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (r *MovieCommandRepository) UpdateMovie(id int, movie models.Movie) (models.Movie, error) {
	if err := r.db.Model(&models.Movie{}).Where("id = ?", id).Updates(movie).Error; err != nil {
		return models.Movie{}, err
	}

	var updatedMovie models.Movie
	if err := r.db.First(&updatedMovie, id).Error; err != nil {
		return models.Movie{}, err
	}

	return updatedMovie, nil
}

func (r *MovieCommandRepository) BulkInsertMovie(movies []models.Movie) error {
	if len(movies) == 0 {
		return nil 
	}
	if err := r.db.Create(&movies).Error; err != nil {
		return err
	}
	return nil
}