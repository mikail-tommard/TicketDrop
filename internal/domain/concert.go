package domain

import "strings"

type Concert struct {
	id          int
	title       string
	description string
}

func NewConcert(id int, title, description string) (*Concert, error) {
	concert := Concert{
		id:          id,
		title:       title,
		description: description,
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
	return nil
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
