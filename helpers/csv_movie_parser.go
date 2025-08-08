package helpers

import (
	"encoding/csv"
	"io"
	"mime/multipart"
	"strconv"

	"challenge-2/models"
)

func ParseCSVToMovies(file multipart.File) ([]models.Movie, error) {
	var movies []models.Movie

	reader := csv.NewReader(file)
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		duration, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}

		movie := models.Movie{
			Title:       record[0],
			Description: record[1],
			Duration:    duration,
			Artists:     record[3],
			Genres:      record[4],
		}
		movies = append(movies, movie)
	}

	return movies, nil
}