package routes

import (
	"challenge-2/config"
	movieController "challenge-2/modules/movies/controllers"
	movieRepository "challenge-2/modules/movies/repositories"
	movieService "challenge-2/modules/movies/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	db := config.ConnectDB()

	movieQueryRepo := movieRepository.NewMovieQuery(db)
	movieCommandRepo := movieRepository.NewMovieCommand(db)
	movieService := movieService.NewMovieService(movieQueryRepo, movieCommandRepo)
	movieController := movieController.NewMovieController(movieService)

	e.GET("/movies", movieController.GetMovies)
	e.POST("/movies", movieController.CreateMovie)
	e.PUT("/movies/:id", movieController.UpdateMovie)
	e.GET("/movies/search", movieController.SearchMovies)
	e.POST("/movies/upload", movieController.UploadMovies)
}
