package db

import (
	"ArtAPI/internal/artobj"
	"context"
	"database/sql"
	"fmt"
)

type ArtRow struct {
	ID                int `db:"id"`
	ObjectID          int `db:"object_id"`
	IsHighlight       bool `db:"is_highlight"`
	AccessionYear     sql.NullString `db:"accession_year"`
	PrimaryImage      sql.NullString `db:"primary_image"`
	Department        sql.NullString `db:"department"`
	Title             sql.NullString `db:"title"`
	ObjectName  	  sql.NullString `db:"object_name"`
	Culture           sql.NullString `db:"culture"`
	Period            sql.NullString `db:"period"`
	ArtistDisplayName sql.NullString `db:"artist_display_name"`
	City              sql.NullString `db:"city"`
	Country           sql.NullString `db:"country"`
}

func convertArtRowtoArtObj(row ArtRow) artobj.ArtObject {
	return artobj.ArtObject{
		ID: row.ID,
		ObjectID: row.ObjectID,
		IsHighlight: row.IsHighlight,
		AccessionYear: row.AccessionYear.String,
		PrimaryImage: row.AccessionYear.String,
		Department: row.Department.String,
		Title: row.Title.String,
		ObjectName: row.ObjectName.String,
		Culture: row.Culture.String,
		Period: row.Period.String,
		City: row.City.String,
		Country: row.Country.String,
	}
}

func (d *Database) GetArt(
	ctx context.Context,
	id int,
) (artobj.ArtObject, error) {
	var artRow ArtRow
	row := d.Client.QueryRowxContext(
		ctx,
		`SELECT *
		FROM art_vault
		WHERE id = $1`,
		id,
	)
	err := row.StructScan(&artRow) 
	if err != nil {
		return artobj.ArtObject{}, fmt.Errorf("error fetching art object by id: %w", err)
	}
	return convertArtRowtoArtObj(artRow), nil 
}