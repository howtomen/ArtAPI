package artobj

import (
	"context"
	"errors"
	"fmt"
)

// using these to pass info on what went wrong
// without exposing internal errors to client.
var (
	ErrFetchingAllArt = errors.New("failed to get all art records")
	ErrFetchingArt = errors.New("failed to fetch art record by id")
	ErrNotImplemented = errors.New("not implemented")
	ErrPostingArt = errors.New("failed to upload art record to database")
	ErrUpdatingRow = errors.New("failed to update record in database")
)

//Store - This interface defines all methods
// that our service needs to operate.
type Store interface {
	GetAllArt(context.Context) ([]ArtObject, error)
	GetArt(context.Context, string) (ArtObject, error)
	PostArt(context.Context, ArtObject) (ArtObject, error)
	UpdateArt(context.Context, string, ArtObject) (ArtObject, error)
	DeleteArt(context.Context, string) (error)
}
// This is an individual Art Object Record in the form of a Struct
type ArtObject struct {
	ID 					string 	`json:"id"` 
	ObjectID 			int 	`json:"object_id"`
	IsHighlight 		bool 	`json:"is_highlight"`
	AccessionYear 		string 	`json:"accession_year"`
	PrimaryImage 		string 	`json:"primary_image"`
	Department 			string	`json:"department"`
	Title 				string 	`json:"title"`
	ObjectName 			string 	`json:"object_name"`
	Culture	 			string 	`json:"culture"`
	Period 				string 	`json:"period"`
	ArtistDisplayName 	string 	`json:"artist_display_name"`
	City 				string 	`json:"city"`
	Country 			string 	`json:"country"`
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

func (s *Service) GetAllArt(ctx context.Context) ([]ArtObject, error) {
	fmt.Println("Getting All Art Objects in vault") 
	art, err := s.Store.GetAllArt(ctx)
	if err != nil{
		fmt.Println(err)
		return []ArtObject{}, ErrFetchingAllArt
	}

	return art, nil 
}

func (s *Service) GetArt(ctx context.Context, id string) (ArtObject, error) {
	fmt.Println("Getting Art Object")
	art, err := s.Store.GetArt(ctx,id)
	if err != nil {
		fmt.Println(err)
		return ArtObject{},ErrFetchingArt 
	}
	
	return art, nil 
}

func (s *Service) UpdateArt(ctx context.Context, id string, art ArtObject) (ArtObject, error) {
	fmt.Println("Updating Art Object")
	response, err := s.Store.UpdateArt(ctx,id,art)
	if err != nil {
		fmt.Println(err)
		return ArtObject{},ErrUpdatingRow
	}
	return response, nil  
}

func (s *Service) DeleteArt(ctx context.Context, id string) error {
	fmt.Println("Deleting Art Object")
	return s.Store.DeleteArt(ctx, id)
}

func (s *Service) PostArt(ctx context.Context, art ArtObject) (ArtObject, error) {
	fmt.Println("Posting Art Object")
	res, err := s.Store.PostArt(ctx,art)
	if err != nil {
		fmt.Println(err)
		return ArtObject{}, ErrPostingArt
	}
	return res, nil
}