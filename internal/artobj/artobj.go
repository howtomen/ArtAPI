package artobj

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingArt = errors.New("failed to fetch art record by id")
	ErrNotImplemented = errors.New("not implemented")
)

//Store - This interface defines all methods
// that our service needs to operate.
type Store interface {
	GetArt(context.Context, int) (ArtObject, error)
}
// This is an individual Art Object Record in the form of a Struct
type ArtObject struct {
	ID 					int `json:"id"` 
	ObjectID 			int `json:"object_id"`
	IsHighlight 		bool `json:"is_highlight"`
	AccessionYear 		string `json:"accession_year"`
	PrimaryImage 		string `json:"primary_image"`
	Department 			string `json:"department"`
	Title 				string `json:"title"`
	ObjectName 			string `json:"object_name"`
	Culture	 			string `json:"culture"`
	Period 				string `json:"period"`
	ArtistDisplayName 	string `json:"artist_display_name"`
	City 				string `json:"city"`
	Country 			string `json:"country"`
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

func (s *Service) GetArt(ctx context.Context, id int) (ArtObject, error) {
	fmt.Println("Getting Art Object")
	art, err := s.Store.GetArt(ctx,id)
	if err != nil {
		fmt.Println(err)
		return ArtObject{},ErrFetchingArt
	}
	
	return art, nil 
}

func (s *Service) UpdateArt(ctx context.Context, art ArtObject) error {
	return ErrNotImplemented
}

func (s *Service) DeleteArt(ctx context.Context, id int) error {
	return ErrNotImplemented
}

func (s *Service) CreateArt(ctx context.Context, art ArtObject) (ArtObject, error) {
	return ArtObject{}, ErrNotImplemented
}