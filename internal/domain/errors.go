package domain

import "errors"

var (
	ErrConcertID     = errors.New("Concert ID invalid")
	ErrConcertTitile = errors.New("Concert Title invalid")
	ErrConcertDesc   = errors.New("Concert Description invalid")

	ErrConcertNotFound = errors.New("Concert not found")
)
