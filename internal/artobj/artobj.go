package artobj

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
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
	ID 					string 	
	ObjectID 			int 	
	IsHighlight 		bool 	
	AccessionYear 		string 	
	PrimaryImage 		string 	
	Department 			string	
	Title 				string 	
	ObjectName 			string 	
	Culture	 			string 	
	Period 				string 	
	ArtistDisplayName 	string 	
	City 				string 	
	Country 			string 
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
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Getting All Art")
	art, err := s.Store.GetAllArt(ctx)
	if err != nil{
		l.Debug().Err(err).Msg("")
		return []ArtObject{}, ErrFetchingAllArt
	}

	return art, nil 
}

func (s *Service) GetArt(ctx context.Context, id string) (ArtObject, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Getting Art Object")
	art, err := s.Store.GetArt(ctx,id)
	if err != nil {
		l.Debug().Err(err).Msg("")
		return ArtObject{},ErrFetchingArt 
	}
	
	return art, nil 
}

func (s *Service) UpdateArt(ctx context.Context, id string, art ArtObject) (ArtObject, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Updating Art Object")
	response, err := s.Store.UpdateArt(ctx,id,art)
	if err != nil {
		l.Debug().Err(err).Msg("")
		return ArtObject{},ErrUpdatingRow
	}
	return response, nil  
}

func (s *Service) DeleteArt(ctx context.Context, id string) error {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Deleting Art Object")
	return s.Store.DeleteArt(ctx, id)
}

func (s *Service) PostArt(ctx context.Context, art ArtObject) (ArtObject, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Posting Art Object")
	res, err := s.Store.PostArt(ctx,art)
	if err != nil {
		l.Debug().Err(err).Msg("")
		return ArtObject{}, ErrPostingArt
	}
	return res, nil
}