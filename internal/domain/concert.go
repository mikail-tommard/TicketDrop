package domain

import "strings"

type Concert struct {
	id          int
	title       string
	description string
	capacity    int
	available   int
}

func NewConcert(id, available, capacity int, title, description string) (*Concert, error) {
	concert := Concert{
		id:          id,
		title:       title,
		description: description,
		capacity:    capacity,
		available:   available,
	}

	err := validate(concert)
	if err != nil {
		return nil, err
	}

	return &concert, nil
}

func validate(concert Concert) error {
	if concert.id <= 0 {
		return ErrConcertID
	}

	concert.title = strings.TrimSpace(concert.title)
	if len(concert.title) <= 0 {
		return ErrConcertTitile
	}

	concert.description = strings.TrimSpace(concert.description)
	if len(concert.description) <= 0 {
		return ErrConcertDesc
	}

	if concert.available < 0 {
		return ErrConcertAvailable
	}

	if concert.capacity < 0 {
		return ErrConcertCapacity
	}
	return nil
}

func FromStorage(id, capacity, available int, title, description string) *Concert {
	return &Concert{
		id:          id,
		capacity:    capacity,
		title:       title,
		description: description,
		available:   available,
	}
}

func (c *Concert) ID() int {
	return c.id
}
func (c *Concert) Title() string {
	return c.title
}
func (c *Concert) Desc() string {
	return c.description
}
func (c *Concert) Available() int {
	return c.available
}
func (c *Concert) Capacity() int {
	return c.capacity
}
