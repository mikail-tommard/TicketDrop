package domain

import "errors"

var (
	ErrConcertID        = errors.New("Concert ID invalid")
	ErrConcertTitile    = errors.New("Concert Title invalid")
	ErrConcertDesc      = errors.New("Concert Description invalid")
	ErrConcertAvailable = errors.New("Concert Available invalid")
	ErrConcertCapacity  = errors.New("Concert Capacity invalid")

	ErrConcertNotFound = errors.New("Concert not found")

	ErrAlreadyBookerd = errors.New("Already booked")
	ErrKey            = errors.New("Key invalid")
	ErrSoldOut        = errors.New("no available seats")
)
