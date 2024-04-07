package main

import (
	db "ArtAPI/internal/DB"
	"ArtAPI/internal/artobj"
	transportHttp "ArtAPI/internal/transport/http"
	// "ArtAPI/routers/middleware"
	// "context"
	"fmt"

	// "log"
	// "net/http"
)
func Run() error {
	fmt.Println("Starting up App Application")
	db, err := db.NewDbConnection()
	if err != nil {
		fmt.Println("Failed to connect to DB")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate db")
		return err 
	}

	artServ := artobj.NewService(db)

	httpHandler := transportHttp.NewHandler(artServ)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}
func main() {
	//open connection to DB
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}