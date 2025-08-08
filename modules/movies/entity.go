package movies

import (
	"challenge-2/helpers"
	"challenge-2/models"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	ID 				string `json:"id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Duration    int    `json:"duration" gorm:"column:duration"`
	Artists     string `json:"artists" gorm:"column:artists"`
	Genres      string `json:"genres" gorm:"column:genres"`
}

type MovieControllerInterface interface {
	GetMovies(c echo.Context) error
	CreateMovie(c echo.Context) error
	UpdateMovie(c echo.Context) error
	SearchMovies(c echo.Context) error 
	UploadMovies(c echo.Context) error
}

type MovieServiceInterface interface {
	GetMovies(p helpers.Pagination) ([]models.Movie, int64, error)
	CreateMovie(movie models.Movie) (models.Movie, error)
	UpdateMovie(id int, data map[string]interface{}) (models.Movie, error)
	SearchMovies(query string, p helpers.Pagination) ([]models.Movie, int64, error)
	UploadMovies(movies []models.Movie) error
}

type MovieQueryInterface interface {
	GetMovies(p helpers.Pagination) ([]models.Movie, int64, error)
	SearchMovies(query string, p helpers.Pagination) ([]models.Movie, int64, error)
}

type MovieCommandInterface interface {
	CreateMovie(movie models.Movie) (models.Movie, error)
	UpdateMovie(id int, data map[string]interface{}) (models.Movie, error)
	BulkInsertMovie(movies []models.Movie) error
}
