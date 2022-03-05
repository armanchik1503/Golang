package domain

import "time"

type Book struct {
	ID          string
	Title       string
	Description string
	Author      string
	Genre       string
	ReleaseDate time.Time
}
