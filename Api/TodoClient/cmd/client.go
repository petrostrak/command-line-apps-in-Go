package cmd

import (
	"errors"
	"time"
)

var (
	ErrConnection      = errors.New("connection error")
	ErrNotFound        = errors.New("not found")
	ErrInvalidResponse = errors.New("invalid server response")
	ErrInvalid         = errors.New("invalid data")
	ErrNotNumber       = errors.New("not a number")
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type response struct {
	Results      []item `json:"results"`
	Date         int    `json:"date"`
	TotalResults int    `json:"total_results"`
}
