package controllers

type CreateMovieRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Duration    int    `json:"duration" validate:"required,min=1"`
	Artists     string `json:"artists" validate:"required"`
	Genres      string `json:"genres" validate:"required"`
}

type UpdateMovieRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Duration    *int    `json:"duration"`
	Artists     *string `json:"artists"`
	Genres      *string `json:"genres"`
}

type UploadMoviesRequest struct {
	Movies []CreateMovieRequest `json:"movies" validate:"required,dive"`
}