package http

import (
	"ArtAPI/internal/artobj"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type ArtService interface {
	GetAllArt(context.Context) ([]artobj.ArtObject, error)
	GetArt(ctx context.Context, id string) (artobj.ArtObject, error)
	PostArt(context.Context, artobj.ArtObject) (artobj.ArtObject, error)
	UpdateArt(ctx context.Context, id string, newArt artobj.ArtObject) (artobj.ArtObject, error)
	DeleteArt(ctx context.Context, id string) (error)
}

type Response struct {
	Message string
}

type PostArtRequest struct {
	ObjectID 			int 	`json:"object_id" validate:"required"`
	IsHighlight 		bool 	`json:"is_highlight" validate:"boolean"`
	AccessionYear 		string 	`json:"accession_year" validate:"numeric"`
	PrimaryImage 		string 	`json:"primary_image"`
	Department 			string	`json:"department"`
	Title 				string 	`json:"title" validate:"required"`
	ObjectName 			string 	`json:"object_name" validate:"required"`
	Culture	 			string 	`json:"culture"`
	Period 				string 	`json:"period"`
	ArtistDisplayName 	string 	`json:"artist_display_name" validate:"required"`
	City 				string 	`json:"city"`
	Country 			string 	`json:"country"`
}

func convertPostRequestToArtObj (req PostArtRequest) (artobj.ArtObject) {
	return artobj.ArtObject{
		ObjectID: req.ObjectID,
		IsHighlight: req.IsHighlight,
		AccessionYear: req.AccessionYear,
		PrimaryImage: req.PrimaryImage,
		Department: req.Department,
		Title: req.Title,
		Culture: req.Culture,
		Period: req.Period,
		ArtistDisplayName: req.ArtistDisplayName,
		City: req.City,
		Country: req.Country,
	}
}

func (h *Handler) GetAllArt(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())
	art, err := h.Service.GetAllArt(r.Context())
	if err != nil {
		l.Info().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(art); err != nil {
		panic(err)
	} 
}

func (h *Handler) GetArt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	l := zerolog.Ctx(r.Context())

	art, err := h.Service.GetArt(r.Context(), id)
	if err != nil {
		l.Info().Err(err).Msg("")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(art); err != nil {
		l.Panic().Err(err).Msg("Error encoding JSON")
	} 
}

func (h *Handler) PostArt(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())
	var art PostArtRequest
	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		l.Info().Err(err).Msg("")
		return
	}

	validate := validator.New()
	err := validate.Struct(art)
	if err != nil {
		l.Info().Err(err).Msg("received invalid JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	convertedArt := convertPostRequestToArtObj(art)
	postedComment, err := h.Service.PostArt(r.Context(), convertedArt)
	if err != nil {
		l.Info().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 

	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		l.Panic().Err(err).Msg("Error encoding JSON")
	} 
}

func (h *Handler) UpdateArt(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())
	var art PostArtRequest
	vars := mux.Vars(r)
	id := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		l.Debug().Err(err).Msg("")
		return
	}

	validate := validator.New()
	err := validate.Struct(art)
	if err != nil {
		l.Info().Err(err).Msg("received invalid JSON")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	convertedArt := convertPostRequestToArtObj(art)
	response, err := h.Service.UpdateArt(r.Context(), id, convertedArt)
	if err != nil {
		l.Info().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		l.Panic().Err(err).Msg("Error encoding JSON")
	} 	
}

func (h *Handler) DeleteArt(w http.ResponseWriter, r *http.Request) {
	l := zerolog.Ctx(r.Context())
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Service.DeleteArt(r.Context(), id)
	if err != nil {
		l.Info().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted record"}); err != nil {
		l.Info().Err(err).Msg("Error encoding JSON")
		return 
	}
}