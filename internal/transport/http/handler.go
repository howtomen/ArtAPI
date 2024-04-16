package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

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
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
	h.Router.Use(TimeoutMiddleware)

	h.Server = &http.Server{
		Addr: "0.0.0.0:8000",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes () {
	h.Router.HandleFunc("/api/v3/art", h.GetAllArt).Methods("GET")
	h.Router.HandleFunc("/api/v3/art/{id}", h.GetArt).Methods("GET")
	h.Router.HandleFunc("/api/v3/art", JWTAuth(h.PostArt)).Methods("POST")
	h.Router.HandleFunc("/api/v3/art/{id}", JWTAuth(h.UpdateArt)).Methods("PUT")
	h.Router.HandleFunc("/api/v3/art/{id}", JWTAuth(h.DeleteArt)).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// this should handle shutdown but by default waits indefinitely
	// we defer cancel so that it doesnt hang 
	defer cancel() 
	err := h.Server.Shutdown(ctx)
	if err != nil {
		log.Print(err)
	}

	log.Println("shut down gracefully")
	return nil 
}
