package main

import (
	db "ArtAPI/internal/DB"
	"ArtAPI/internal/artobj"
	// "ArtAPI/routers"
	// "ArtAPI/routers/middleware"
	"context"
	"fmt"
	// "log"
	// "net/http"
)
func Run() error {
	fmt.Println("This is eventually going to run the Application.")
	db, err := db.NewDbConnection()
	if err != nil {
		fmt.Println("Failed to connect to DB")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	artServ := artobj.NewService(db)
	fmt.Println(artServ.GetArt(
		context.Background(),
		1,
	))
	return nil
}
func main() {
	//open connection to DB
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}