package controllers

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Duration    int    `json:"duration" validate:"required,min=1"`
	Artists     string `json:"artists" validate:"required"`
	Genres      string `json:"genres" validate:"required"`
}

type UpdateMovieRequest struct {
	Title       string `json:"title" validate:"nullable"`
	Description string `json:"description" validate:"nullable"`
	Duration    int    `json:"duration" validate:"nullable,min=1"`
	Artists     string `json:"artists" validate:"nullable"`
	Genres      string `json:"genres" validate:"nullable"`
}

type UploadMoviesRequest struct {
	Movies []CreateMovieRequest `json:"movies" validate:"required,dive"`
}