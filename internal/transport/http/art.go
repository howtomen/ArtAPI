package http

import (
	"ArtAPI/internal/artobj"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArtService interface {
	GetAllArt(context.Context) ([]artobj.ArtObject, error)
	GetArt(ctx context.Context, id int) (artobj.ArtObject, error)
	PostArt(context.Context, artobj.ArtObject) (artobj.ArtObject, error)
	UpdateArt(ctx context.Context, id int, newArt artobj.ArtObject) (artobj.ArtObject, error)
	DeleteArt(ctx context.Context, id int) (error)
}

type Response struct {
	Message string
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	var art artobj.ArtObject
	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		log.Print(err)
		return
	}
	art, err := h.Service.PostArt(r.Context(), art)
	if err != nil {
		log.Print(err)
		return
	} 

	if err := json.NewEncoder(w).Encode(art); err != nil {
		panic(err)
	} 
}

func (h *Handler) UpdateArt(w http.ResponseWriter, r *http.Request) {
	var art artobj.ArtObject
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&art); err != nil {
		log.Print(err)
		return
	}

	response, err := h.Service.UpdateArt(r.Context(), id, art)
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteArt(r.Context(), id)
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