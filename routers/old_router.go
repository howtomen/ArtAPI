package router

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	
	router.HandleFunc("/objects",getObjects(db)).Methods("GET")
	router.HandleFunc("/objects/{id}", getObject(db)).Methods("GET")
	router.HandleFunc("/objects", createObject(db)).Methods("POST")
	router.HandleFunc("/objects/{id}", updateObject(db)).Methods("PUT")
	router.HandleFunc("/objects/{id}", deleteObject(db)).Methods("DELETE")

	return router
}