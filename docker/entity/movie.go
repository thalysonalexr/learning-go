package entity

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Poster      string    `json:"poster"`
	Overview    string    `json:"overview"`
	ReleaseNote time.Time `json:"release_note"`
	Genres      []string  `json:"genres"`
}
