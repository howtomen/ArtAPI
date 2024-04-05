package main

import (
	database "ArtAPI/DB"
	"ArtAPI/routers"
	"ArtAPI/routers/middleware"
	"fmt"
	"log"
	"net/http"
)
func Run() error {
	fmt.Println("This is eventually going to run the Application.")
	return nil
}
func main() {
	//open connection to DB
	db := database.NewDbConnection()
	defer db.Close()

	//create router
	router := router.NewRouter(db)
	log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(router)))
}