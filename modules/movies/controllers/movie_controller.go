package controllers

import (
	"challenge-2/helpers"
	"challenge-2/models"
	"challenge-2/modules/movies"
	"encoding/csv"
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
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Gagal mengambil data film"))
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

	return c.JSON(http.StatusOK, helpers.SuccessPaginationResponse("Berhasil mengambil data film", responseData, paginationInfo))
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
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Gagal membuat film"))
	}
	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Film berhasil dibuat", GetMovieResponse{
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

	if req.Title == "" && req.Description == "" && req.Duration <= 0 && req.Artists == "" && req.Genres == "" {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("No fields to update"))
	}

	modelMovie := models.Movie{
		Title:       req.Title,
		Description: req.Description,
		Duration:    req.Duration,
		Artists:     req.Artists,
		Genres:      req.Genres,
	}
	updatedMovie, err := mc.s.UpdateMovie(id, modelMovie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Gagal memperbarui film"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Film berhasil diperbarui", GetMovieResponse{
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
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Gagal mencari film"))
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

	return c.JSON(http.StatusOK, helpers.SuccessPaginationResponse("Berhasil mencari film", responseData, paginationInfo))
}

func (mc *MovieController) UploadMovies(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Gagal membaca file CSV"))
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Gagal membuka file"))
	}
	defer src.Close()

	reader := csv.NewReader(src)
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Gagal membaca isi file CSV"))
	}
	if len(records) < 2 {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Data CSV kosong atau tidak lengkap"))
	}

	var movies []models.Movie
	for i, row := range records {
		if i == 0 {
			continue
		}

		if len(row) < 5 {
			continue
		}

		duration, err := strconv.Atoi(row[2])
		if err != nil {
			continue
		}

		movie := models.Movie{
			Title:       row[0],
			Description: row[1],
			Duration:    duration,
			Artists:     row[3],
			Genres:      row[4],
		}
		movies = append(movies, movie)
	}

	if len(movies) == 0 {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Tidak ada data movie yang valid untuk disimpan"))
	}

	err = mc.s.UploadMovies(movies)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Gagal menyimpan data movie"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Berhasil mengunggah dan menyimpan data movie", nil))
}