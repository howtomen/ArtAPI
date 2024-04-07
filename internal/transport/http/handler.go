package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ArtService interface{}

type Handler struct {
	Router  *mux.Router
	Service ArtService
	Server *http.Server
}

func NewHandler(service ArtService) *Handler {
	h := &Handler{
		Service: service, 
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8000",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes () {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil 
}