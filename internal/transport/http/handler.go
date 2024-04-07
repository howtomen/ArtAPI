package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")
	return nil 
}