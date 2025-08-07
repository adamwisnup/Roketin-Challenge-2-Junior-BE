package controllers

type GetMovieResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genres      string `json:"genres"`
	Artists     string `json:"artists"`
	Duration    int    `json:"duration"`
}