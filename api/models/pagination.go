package models

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidPageParameter  = errors.New("invalid page parameter")
	ErrInvalidLimitParameter = errors.New("invalid limit parameter")
)

type Pagination struct {
	PageIndex int `json:"pageIndex"`
	Limit     int `json:"limit"`
}

func NewPagination(pageStr, limitStr string) *Pagination {
	pageIndex, err := strconv.Atoi(pageStr)
	if err != nil || pageIndex < 1 {
		pageIndex = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	return &Pagination{
		PageIndex: pageIndex,
		Limit:     limit,
	}
}

type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"totalPages"`
	PageIndex  int   `json:"pageIndex"`
	Limit      int   `json:"limit"`
}

func MapPaginatedResult[T any, U any](result *PaginatedResponse[T], mapper func(T) U) *PaginatedResponse[U] {
	newData := make([]U, len(result.Data))
	for i, item := range result.Data {
		newData[i] = mapper(item)
	}

	return &PaginatedResponse[U]{
		Data:       newData,
		Total:      result.Total,
		TotalPages: result.TotalPages,
		PageIndex:  result.PageIndex,
		Limit:      result.Limit,
	}
}
