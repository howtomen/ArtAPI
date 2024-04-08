package db

import (
	"ArtAPI/internal/artobj"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type ArtRow struct {
	ID                string `db:"id"`
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
		ID: 				row.ID,
		ObjectID: 			row.ObjectID,
		IsHighlight: 		row.IsHighlight,
		AccessionYear: 		row.AccessionYear.String,
		PrimaryImage: 		row.PrimaryImage.String,
		Department: 		row.Department.String,
		Title: 				row.Title.String,
		ObjectName: 		row.ObjectName.String,
		Culture: 			row.Culture.String,
		Period: 			row.Period.String,
		ArtistDisplayName: 	row.ArtistDisplayName.String,
		City: 				row.City.String,
		Country: 			row.Country.String,
	}
}

func convertObjToRow(art artobj.ArtObject) ArtRow {
	row := ArtRow{
		ID: 				art.ID,
		ObjectID: 			art.ObjectID,
		IsHighlight: 		art.IsHighlight,
		AccessionYear: 		sql.NullString{String:art.AccessionYear, Valid: true},
		PrimaryImage:		sql.NullString{String:art.PrimaryImage, Valid: true},
		Department: 		sql.NullString{String:art.Department, Valid: true},
		Title: 				sql.NullString{String:art.Title, Valid: true},
		ObjectName: 		sql.NullString{String:art.ObjectName, Valid: true},
		Culture: 			sql.NullString{String:art.Culture, Valid: true},
		Period: 			sql.NullString{String:art.Period, Valid: true},
		ArtistDisplayName:	sql.NullString{String:art.ArtistDisplayName, Valid: true},
		City: 				sql.NullString{String:art.City, Valid: true},
		Country: 			sql.NullString{String:art.Country, Valid: true},

	}
	return row
}

func (d *Database) GetAllArt(ctx context.Context) ([]artobj.ArtObject, error) {
	var allArt []artobj.ArtObject
	rows, err := d.Client.Queryx("SELECT * FROM art_vault")
	if err != nil {
		return allArt, fmt.Errorf("error fetching all art objects: %w", err)
	}

	for rows.Next() {
		row := ArtRow{}
		if err := rows.StructScan(&row); err != nil {
			return allArt, fmt.Errorf("error scanning row: %w", err)
		}
		allArt = append(allArt, convertArtRowtoArtObj(row))
	}

	if err := rows.Close(); err != nil {
		return allArt, fmt.Errorf("failed to close rows: %w", err)
	}	
	return allArt, nil 
} 

func (d *Database) GetArt(ctx context.Context, id string,) (artobj.ArtObject, error) {
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

func (d *Database) PostArt(ctx context.Context, art artobj.ArtObject) (artobj.ArtObject, error) {
	art.ID = uuid.NewString() //generate uuid string 
	postRow := convertObjToRow(art)

	_, err := d.Client.NamedExecContext(
		ctx,
		"INSERT INTO art_vault (id,object_id,is_highlight,accession_year,primary_image,department,title,object_name,culture,period,artist_display_name,city,country) VALUES (:id,:object_id,:is_highlight,:accession_year,:primary_image,:department,:title,:object_name,:culture,:period,:artist_display_name,:city,:country);",
		postRow,
	)
	if err != nil {
		return artobj.ArtObject{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	return art, nil 
}

func (d *Database) DeleteArt(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM art_vault WHERE id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete art record from database: %w", err)
	}

	return nil
}

func (d *Database) UpdateArt(ctx context.Context, id string, art artobj.ArtObject) (artobj.ArtObject, error) {
	art.ID = id
	artRow := convertObjToRow(art)
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE art_vault SET object_id = :object_id,is_highlight = :is_highlight,accession_year = :accession_year,primary_image = :primary_image,department = :department,title = :title,object_name = :object_name,culture = :culture,period = :period,artist_display_name = :artist_display_name,city = :city,country = :country WHERE id=:id;`,
		artRow, 
	)

	if err != nil {
		return artobj.ArtObject{}, fmt.Errorf("failed to update comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return artobj.ArtObject{}, fmt.Errorf("failed to close rows: %w", err)
	}
	return art, nil 
}