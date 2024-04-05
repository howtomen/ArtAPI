package artobj

import (
	"context"
	"fmt"
)

type Store interface {
	GetArt(context.Context, string) (ArtObject, error)
}
// This is an individual Art Object Record in the form of a Struct
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

//Service - the struct on which all of our logic will be built ontop of
type Service struct{
	Store Store
}

//NewService - returns pointer to new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetArt(ctx context.Context, id string) (ArtObject, error) {
	fmt.Println("Getting Art Object")
	art, err := s.Store.GetArt(ctx,id)
	if err != nil {
		fmt.Println(err)
		return ArtObject{},err
	}
	
	return art, nil 
}
