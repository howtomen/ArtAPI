package models

type ArtObject struct {
	ID int `json:"id" db:"id"` 
	ObjectID int `json:"object_id" db:"object_id"`
	IsHighlight bool `json:"is_highlight" db:"is_highlight"`
	AccessionYear string `json:"accession_year" db:"accession_year"`
	PrimaryImage string `json:"primary_image" db:"primary_image"`
	Department string `json:"department" db:"department"`
	Title string `json:"title" db:"title"`
	ObjectName string `json:"object_name" db:"object_name"`
	Culture string `json:"culture" db:"culture"`
	Period string `json:"period" db:"period"`
	ArtistDisplayName string `json:"artist_display_name" db:"artist_display_name"`
	City string `json:"city" db:"city"`
	Country string `json:"country" db:"country"`
}
