package controllers

import (
	"challenge-2/helpers"
	"challenge-2/models"
	"challenge-2/modules/movies"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MovieController struct {
	s movies.MovieServiceInterface
}

func NewMovieController(s movies.MovieServiceInterface) *MovieController {
	return &MovieController{s: s}
}

func (mc *MovieController) GetMovies(c echo.Context) error {
	pagination := helpers.GetPaginationParams(c)

	movies, totalItems, err := mc.s.GetMovies(pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to retrieve movies"))
	}

	totalPages := int((totalItems + int64(pagination.Limit) - 1) / int64(pagination.Limit))

	var responseData []GetMovieResponse
	for _, m := range movies {
		responseData = append(responseData, GetMovieResponse{
			ID:          m.ID,
			Title:       m.Title,
			Description: m.Description,
			Genres:      m.Genres,
			Artists:     m.Artists,
			Duration:    m.Duration,
		})
	}

	paginationInfo := helpers.PaginationInfo{
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	return c.JSON(http.StatusOK, helpers.SuccessPaginationResponse("Movies retrieved successfully", responseData, paginationInfo))
}

func (mc *MovieController) CreateMovie(c echo.Context) error {
	var req CreateMovieRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid request body"))
	}

	if req.Title == "" || req.Description == "" || req.Duration <= 0 {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid movie data"))
	}

	modelMovie := models.Movie{
		Title:       req.Title,
		Description: req.Description,
		Duration:    req.Duration,
		Artists:     req.Artists,
		Genres:      req.Genres,
	}

	createdMovie, err := mc.s.CreateMovie(modelMovie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to create movie"))
	}
	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Movie created successfully", GetMovieResponse{
		ID:          createdMovie.ID,
		Title:       createdMovie.Title,
		Description: createdMovie.Description,
		Genres:      createdMovie.Genres,
		Artists:     createdMovie.Artists,
		Duration:    createdMovie.Duration,
	}))
}

func (mc *MovieController) UpdateMovie(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid movie ID"))
	}

	var req UpdateMovieRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid request body"))
	}

	if req.Title == nil && req.Description == nil && req.Duration == nil && req.Artists == nil && req.Genres == nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("No fields to update"))
	}

	updateData := make(map[string]interface{})

	if req.Title != nil {
		updateData["title"] = *req.Title
	}
	if req.Description != nil {
		updateData["description"] = *req.Description
	}
	if req.Duration != nil {
		updateData["duration"] = *req.Duration
	}
	if req.Artists != nil {
		updateData["artists"] = *req.Artists
	}
	if req.Genres != nil {
		updateData["genres"] = *req.Genres
	}

	updatedMovie, err := mc.s.UpdateMovie(id, updateData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to update movie"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Movie updated successfully", GetMovieResponse{
		ID:          updatedMovie.ID,
		Title:       updatedMovie.Title,
		Description: updatedMovie.Description,
		Genres:      updatedMovie.Genres,
		Artists:     updatedMovie.Artists,
		Duration:    updatedMovie.Duration,
	}))
}

func (mc *MovieController) SearchMovies(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Search query cannot be empty"))
	}

	pagination := helpers.GetPaginationParams(c)
	movies, totalItems, err := mc.s.SearchMovies(query, pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to search movies"))
	}

	totalPages := int((totalItems + int64(pagination.Limit) - 1) / int64(pagination.Limit))

	var responseData []GetMovieResponse
	for _, m := range movies {
		responseData = append(responseData, GetMovieResponse{
			ID:          m.ID,
			Title:       m.Title,
			Description: m.Description,
			Genres:      m.Genres,
			Artists:     m.Artists,
			Duration:    m.Duration,
		})
	}

	paginationInfo := helpers.PaginationInfo{
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}

	return c.JSON(http.StatusOK, helpers.SuccessPaginationResponse("Movies search successful", responseData, paginationInfo))
}

func (mc *MovieController) UploadMovies(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Failed to read CSV file"))
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to open CSV file"))
	}
	defer src.Close()

	movies, err := helpers.ParseCSVToMovies(src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error()))
	}

	err = mc.s.UploadMovies(movies)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to save movie data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Movies uploaded and saved successfully", nil))
}