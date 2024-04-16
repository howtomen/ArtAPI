package main

import (
	db "ArtAPI/internal/DB"
	"ArtAPI/internal/artobj"
	logger "ArtAPI/util/logging"
	transportHttp "ArtAPI/internal/transport/http"
)
func Run() error {
	l := logger.GetLogger()
	l.Info().Msg("Starting up App Application")
	db, err := db.NewDbConnection()
	if err != nil {
		l.Info().Msg("Unable to connect to DB")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		l.Info().Msg("failed to migrate db")
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
	l := logger.GetLogger()
	if err := Run(); err != nil {
		l.Fatal().Err(err).Msg("Art API has encountered an issue and is shutting down.")
	}
}