package models

type Movie struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Year        int32    `json:"year"`
	Country     string   `json:"country"`
	Genres      []string `json:"genres"`
	PosterURL   string   `json:"poster_url"`
	Rating      float32  `json:"rating"`
}
