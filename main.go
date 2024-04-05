package main

import (
	database "ArtAPI/DB"
	"ArtAPI/routers"
	"ArtAPI/routers/middleware"
	"log"
	"net/http"
)

func main() {
	//open connection to DB
	db := database.NewDbConnection()

	//create router
	router := router.NewRouter(db)
	log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(router)))
}