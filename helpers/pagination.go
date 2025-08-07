package helpers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func GetPaginationParams(c echo.Context) Pagination {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return Pagination{
		Page:  page,
		Limit: limit,
	}
}

func Paginate(p Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.Limit
		return db.Offset(offset).Limit(p.Limit)
	}
}
