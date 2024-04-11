package http

import (
	"ArtAPI/internal/artobj"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/go-playground/validator/v10"
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
	IsHighlight 		bool 	`json:"is_highlight" validate:"required"`
	AccessionYear 		string 	`json:"accession_year" validate:"required"`
	PrimaryImage 		string 	`json:"primary_image"`
	Department 			string	`json:"department" validate:"required"`
	Title 				string 	`json:"title" validate:"required"`
	ObjectName 			string 	`json:"object_name" validate:"required"`
	Culture	 			string 	`json:"culture" validate:"required"`
	Period 				string 	`json:"period" validate:"required"`
	ArtistDisplayName 	string 	`json:"artist_display_name" validate:"required"`
	City 				string 	`json:"city" validate:"required"`
	Country 			string 	`json:"country" validate:"required"`
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
	art, err := h.Service.GetAllArt(r.Context())
	if err != nil {
		log.Print(err)
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

	art, err := h.Service.GetArt(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(art); err != nil {
		panic(err)
	} 
}

func (h *Handler) PostArt(w http.ResponseWriter, r *http.Request) {
	var art PostArtRequest
	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		log.Print(err)
		return
	}

	validate := validator.New()
	err := validate.Struct(art)
	if err != nil {
		http.Error(w, "not a valid art object", http.StatusBadRequest)
		return 
	}

	convertedArt := convertPostRequestToArtObj(art)
	postedComment, err := h.Service.PostArt(r.Context(), convertedArt)
	if err != nil {
		log.Print(err)
		return
	} 

	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		panic(err)
	} 
}

func (h *Handler) UpdateArt(w http.ResponseWriter, r *http.Request) {
	var art PostArtRequest
	vars := mux.Vars(r)
	id := vars["id"]

	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		log.Print(err)
		return
	}
	
	validate := validator.New()
	err := validate.Struct(art)
	if err != nil {
		http.Error(w, "not a valid art object", http.StatusBadRequest)
		return 
	}

	convertedArt := convertPostRequestToArtObj(art)
	response, err := h.Service.UpdateArt(r.Context(), id, convertedArt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	} 	
}

func (h *Handler) DeleteArt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Service.DeleteArt(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted record"}); err != nil {
		log.Print(err)
		return 
	}
}