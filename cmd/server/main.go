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
	objs, _ := artServ.GetAllArt(
		context.Background(),
	)
	fmt.Println(objs)
	return nil
}
func main() {
	//open connection to DB
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}