package model

import "errors"

var (
	ErrInvalidRating = errors.New("rating must be greater than 0")
	ErrDBInternal    = errors.New("db internal err")
	ErrNotFound      = errors.New("entity not found")
)
